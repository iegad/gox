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

var upgrader = websocket.Upgrader{}

// 引擎信息
type EngineInfo struct {
	Name string `json:"name"`
}

// 引擎接口
type IEngine interface {
	Info() *EngineInfo                    // 引擎信息
	OnConnected(sess *Sess) error         // 会话连接事件
	OnDisconnected(sess *Sess)            // 会话连接断开事件
	OnData(sess *Sess, data []byte) error // 接收数据事件
	OnRun(iosvc *IOService) error         // 服务启动事件
	OnStopped(iosvc *IOService)           // 服务停止事件
}

// IO服务信息
type IOServiceInfo struct {
	TcpHost      string      `json:"tcp_host,omitempty"`   // TCP监听地址
	WsHost       string      `json:"ws_host,omitempty"`    // websocket监听地址
	TcpConnCount int32       `json:"tcp_connection_count"` // Tcp连接数
	WsConnCount  int32       `json:"ws_connection_count"`  // websocket连接数
	MaxConn      int32       `json:"max_conn"`             // 最大连接数
	Running      bool        `json:"running"`              // 运行状态
	Timeout      int32       `json:"timeout(s)"`           // 超时
	Engine       *EngineInfo `json:"engine"`               // 服务引擎信息
}

func (this_ *IOServiceInfo) String() string {
	jstr, _ := json.Marshal(this_)
	return string(jstr)
}

// IO服务配置
type IOServiceConfig struct {
	TcpEndpoint *string `yaml:"tcp_endpoint,omitempty"`
	WsEndpoint  *string `yaml:"ws_endpoint,omitempty"`
	MaxConn     int32   `yaml:"max_conn,omitempty"`
	Timeout     int32   `yaml:"timeout(s),omitempty"`
}

// IO服务
type IOService struct {
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

	if cfg.Timeout < 0 {
		log.Fatal("timeout is less than 0")
	}

	this_ := &IOService{
		tcpAddr:      tcpAddr,
		wsAddr:       wsAddr,
		tcpListener:  nil,
		wsListener:   nil,
		maxConn:      maxConn,
		tcpConnCount: 0,
		wsConnCount:  0,
		timeout:      time.Duration(cfg.Timeout) * time.Second,
		engine:       engine,
	}

	if this_.wsAddr != nil {
		http.HandleFunc("/ws", this_.wsHandler)
	}

	return this_, nil
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
		Running:      this_.tcpListener != nil || this_.wsListener != nil,
		Timeout:      int32(this_.timeout / time.Second),
		Engine:       this_.engine.Info(),
	}
}

func (this_ *IOService) Engine() IEngine {
	return this_.engine
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

	if this_.tcpListener != nil || this_.wsListener != nil {
		this_.engine.OnStopped(this_)
	}

	if this_.tcpListener != nil {
		this_.tcpListener.Close()
		this_.tcpListener = nil
	}

	if this_.wsListener != nil {
		this_.wsListener.Close()
		this_.wsListener = nil
	}

	this_.wg.Wait()
}

func (this_ *IOService) Wait() {
	this_.wg.Wait()
}

func (this_ *IOService) Run() (err error) {
	this_.Shutdown()

	if this_.tcpAddr == nil && this_.wsAddr == nil {
		err = ErrNoPoints
		return
	}

	if this_.tcpAddr != nil {
		this_.tcpListener, err = net.ListenTCP("tcp", this_.tcpAddr.(*net.TCPAddr))
		if err != nil {
			return
		}
	}

	if this_.wsAddr != nil {
		this_.wsListener, err = net.ListenTCP("tcp", this_.wsAddr.(*net.TCPAddr))
		if err != nil {
			return
		}
	}

	if this_.tcpListener != nil {
		this_.wg.Add(1)
		go this_.tcpRun(&this_.wg)
	}

	if this_.wsListener != nil {
		this_.wg.Add(1)
		go this_.wsRun(&this_.wg)
	}

	err = this_.engine.OnRun(this_)
	if err != nil {
		this_.Shutdown()
	}

	return
}

func (this_ *IOService) tcpRun(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		conn, err := this_.tcpListener.AcceptTCP()
		if err != nil {
			log.Error(err)
			if netIsClosedErr(err) {
				break
			}
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
		atomic.AddInt32(&this_.tcpConnCount, -1)
		conn.Close()
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
		atomic.AddInt32(&this_.wsConnCount, -1)
		conn.Close()
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
			if netIsClosedErr(err) {
				break
			}
		}

		err = this_.engine.OnData(sess, rbuf)
		if err != nil {
			log.Error(err)
			break
		}
	}
}
