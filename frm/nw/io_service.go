package nw

import (
	"encoding/json"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/iegad/gox/frm/log"
)

var (
	upgrader = websocket.Upgrader{}
)

type EngineInfo struct {
	Name string `json:"name"`
}

type IEngine interface {
	Info() *EngineInfo

	OnConnected(sess *Sess) error
	OnDisconnected(sess *Sess)
	OnData(sess *Sess, data []byte) error

	OnRun(iosvc *IOService) error
	OnStopped(iosvc *IOService)
	OnShutdown(iosvc *IOService)
}

type IOServiceInfo struct {
	TcpHost      string      `json:"tcp_host,omitempty"`
	WsHost       string      `json:"ws_host,omitempty"`
	TcpConnCount int32       `json:"tcp_connection_count"`
	WsConnCount  int32       `json:"ws_connection_count"`
	MaxConn      int32       `json:"max_conn"`
	Running      bool        `json:"running"`
	Timeout      int32       `json:"timeout(s)"`
	Engine       *EngineInfo `json:"engine"`
}

func (this_ *IOServiceInfo) String() string {
	jstr, _ := json.Marshal(this_)
	return string(jstr)
}

type IOServiceConfig struct {
	TcpEndpoint *string `yaml:"tcp_endpoint,omitempty"`
	WsEndpoint  *string `yaml:"ws_endpoint,omitempty"`
	MaxConn     int32   `yaml:"max_conn,omitempty"`
	Timeout     int32   `yaml:"timeout(s),omitempty"`
}

type IOService struct {
	running      bool
	tcpAddr      net.Addr
	wsAddr       net.Addr
	tcpListener  *net.TCPListener
	wsListener   *net.TCPListener
	maxConn      int32
	tcpConnCount int32
	wsConnCount  int32
	timeout      time.Duration
	engine       IEngine
	stMtx        sync.Mutex
	wg           sync.WaitGroup
}

func NewIOService(cfg *IOServiceConfig, engine IEngine) (*IOService, error) {
	if cfg == nil {
		return nil, ErrConfigIsNil
	}

	if cfg.TcpEndpoint == nil && cfg.WsEndpoint == nil {
		return nil, ErrNoPoints
	}

	var (
		tcpAddr net.Addr
		wsAddr  net.Addr
		err     error
	)

	if cfg.TcpEndpoint != nil {
		tcpAddr, err = net.ResolveTCPAddr("tcp", *cfg.TcpEndpoint)
		if err != nil {
			return nil, ErrTcpEPIsInvalid
		}
	}

	if cfg.WsEndpoint != nil {
		wsAddr, err = net.ResolveTCPAddr("tcp", *cfg.WsEndpoint)
		if err != nil {
			return nil, ErrWsEPIsInvalid
		}
	}

	maxConn := cfg.MaxConn
	if maxConn < 0 {
		maxConn = 0
	}

	timeout := cfg.Timeout
	if timeout <= 0 || timeout >= MAX_TIMEOUT {
		timeout = DEFAULT_TIMEOUT
	}

	return &IOService{
		running:      false,
		tcpAddr:      tcpAddr,
		wsAddr:       wsAddr,
		tcpListener:  nil,
		wsListener:   nil,
		maxConn:      maxConn,
		tcpConnCount: 0,
		wsConnCount:  0,
		timeout:      time.Duration(timeout * int32(time.Second)),
		engine:       engine,
	}, nil
}

func (this_ *IOService) Info() *IOServiceInfo {

	var (
		tcpHost = ""
		wsHost  = ""
	)

	if this_.tcpAddr != nil {
		tcpHost = this_.tcpAddr.String()
	}

	if this_.wsAddr != nil {
		wsHost = this_.wsAddr.String()
	}

	return &IOServiceInfo{
		TcpHost:      tcpHost,
		WsHost:       wsHost,
		TcpConnCount: this_.tcpConnCount,
		WsConnCount:  this_.wsConnCount,
		MaxConn:      this_.maxConn,
		Running:      this_.running,
		Timeout:      int32(this_.timeout / time.Second),
		Engine:       this_.engine.Info(),
	}
}

func (this_ *IOService) TcpAddr() net.Addr {
	if this_.tcpAddr != nil {
		return this_.tcpAddr
	}

	return nil
}

func (this_ *IOService) WsAddr() net.Addr {
	if this_.wsAddr != nil {
		return this_.wsAddr
	}

	return nil
}

func (this_ *IOService) Shutdown() {
	this_.stMtx.Lock()
	defer this_.stMtx.Unlock()

	if this_.tcpListener != nil {
		this_.tcpListener.Close()
		this_.tcpListener = nil
	}

	if this_.wsListener != nil {
		this_.wsListener.Close()
		this_.wsListener = nil
	}

	this_.engine.OnShutdown(this_)
}

func (this_ *IOService) Run() error {
	this_.Shutdown()

	var err error

	if this_.tcpAddr == nil && this_.wsAddr == nil {
		return ErrNoPoints
	}

	if this_.tcpAddr != nil {
		this_.tcpListener, err = net.ListenTCP("tcp", this_.tcpAddr.(*net.TCPAddr))
		if err != nil {
			return err
		}
	}

	if this_.wsAddr != nil {
		this_.wsListener, err = net.ListenTCP("tcp", this_.wsAddr.(*net.TCPAddr))
		if err != nil {
			return err
		}
	}

	this_.running = true
	err = this_.engine.OnRun(this_)
	if err != nil {
		return err
	}

	if this_.tcpListener != nil {
		this_.wg.Add(1)
		go this_.tcpRun(&this_.wg)
	}

	if this_.wsListener != nil {
		this_.wg.Add(1)
		go this_.wsRun(&this_.wg)
	}

	this_.wg.Wait()
	this_.running = false
	this_.engine.OnStopped(this_)
	return nil
}

func (this_ *IOService) tcpRun(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		conn, err := this_.tcpListener.AcceptTCP()
		if err != nil {
			log.Error(err)
			break
		}

		if this_.maxConn > 0 && atomic.LoadInt32(&this_.tcpConnCount)+atomic.LoadInt32(&this_.wsConnCount) >= this_.maxConn {
			conn.Close()
			continue
		}

		atomic.AddInt32(&this_.tcpConnCount, 1)
		wg.Add(1)
		go this_.tcpConnHandle(conn, wg)
	}
}

func (this_ *IOService) tcpConnHandle(conn *net.TCPConn, wg *sync.WaitGroup) {
	var (
		err  error
		sess = NewTcpSession(conn, this_.timeout)
		rbuf = []byte{}
	)

	defer func() {
		this_.engine.OnDisconnected(sess)
		conn.Close()
		atomic.AddInt32(&this_.tcpConnCount, -1)
		wg.Done()
	}()

	err = this_.engine.OnConnected(sess)
	if err != nil {
		log.Error(err)
		return
	}

	for {
		if this_.timeout > 0 {
			err = conn.SetReadDeadline(time.Now().Add(this_.timeout))
			if err != nil {
				log.Error(err)
				break
			}
		}

		rbuf, err = sess.tcpRead()
		if err != nil {
			log.Error(err)
			break
		}

		err = this_.engine.OnData(sess, rbuf)
		if err != nil {
			log.Error(err)
			break
		}
	}
}

func (this_ *IOService) wsRun(wg *sync.WaitGroup) {
	defer wg.Done()

	http.HandleFunc("/ws", this_.wsHandler)
	err := http.Serve(this_.wsListener, nil)
	if err != nil {
		log.Error(err)
	}
}

func (this_ *IOService) wsHandler(w http.ResponseWriter, r *http.Request) {
	if this_.maxConn > 0 && atomic.LoadInt32(&this_.tcpConnCount)+atomic.LoadInt32(&this_.wsConnCount) >= this_.maxConn {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}

	atomic.AddInt32(&this_.wsConnCount, 1)
	this_.wg.Add(1)
	go this_.wsConnHandle(conn, &this_.wg)
}

func (this_ *IOService) wsConnHandle(conn *websocket.Conn, wg *sync.WaitGroup) {
	var (
		err  error
		sess = NewWsSession(conn, this_.timeout)
		rbuf = []byte{}
	)

	defer func() {
		this_.engine.OnDisconnected(sess)
		conn.Close()
		atomic.AddInt32(&this_.wsConnCount, -1)
		wg.Done()
	}()

	err = this_.engine.OnConnected(sess)
	if err != nil {
		log.Error(err)
		return
	}

	for {
		if this_.timeout > 0 {
			err = conn.SetReadDeadline(time.Now().Add(this_.timeout))
			if err != nil {
				log.Error(err)
				break
			}
		}

		_, rbuf, err = conn.ReadMessage()
		if err != nil {
			log.Error(err)
			break
		}

		err = this_.engine.OnData(sess, rbuf)
		if err != nil {
			log.Error(err)
			break
		}
	}
}
