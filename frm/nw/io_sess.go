package nw

import (
	"encoding/binary"
	"net"
	"time"

	"github.com/gorilla/websocket"
	"github.com/iegad/gox/frm/log"
)

type Sess struct {
	tcpConn  *net.TCPConn
	wsConn   *websocket.Conn
	timeout  time.Duration
	nHead    uint32
	rbuf     []byte
	UserData interface{}
}

func NewTcpSession(conn *net.TCPConn, timeout time.Duration) *Sess {
	if conn == nil {
		log.Fatal("conn is nil")
	}

	return &Sess{
		tcpConn:  conn,
		wsConn:   nil,
		timeout:  timeout,
		nHead:    0,
		rbuf:     []byte{},
		UserData: nil,
	}
}

func NewWsSession(conn *websocket.Conn, timeout time.Duration) *Sess {
	if conn == nil {
		log.Fatal("conn is nil")
	}

	return &Sess{
		tcpConn:  nil,
		wsConn:   conn,
		timeout:  timeout,
		nHead:    0,
		rbuf:     []byte{},
		UserData: nil,
	}
}

func (this_ *Sess) Shutdown() {
	if this_.tcpConn != nil {
		this_.tcpConn.Close()
	}
}

func (this_ *Sess) Write(data []byte) (int, error) {
	if this_.tcpConn != nil {
		return this_.tcpWrite(data)
	}

	return this_.wsWrite(data)
}

func (this_ *Sess) Read() ([]byte, error) {
	if this_.tcpConn != nil {
		return this_.tcpRead()
	}

	return this_.wsRead()
}

func (this_ *Sess) RemoteAddr() net.Addr {
	if this_.tcpConn != nil {
		return this_.tcpConn.RemoteAddr()
	}

	return this_.wsConn.RemoteAddr()
}

func (this_ *Sess) LocalAddr() net.Addr {
	if this_.tcpConn != nil {
		return this_.tcpConn.LocalAddr()
	}

	return this_.wsConn.LocalAddr()
}

func (this_ *Sess) tcpRead() ([]byte, error) {
	if this_.tcpConn == nil {
		log.Fatal("Sess.tcpConn is nil")
	}

	var (
		rbuf  = make([]byte, DEFAULT_RBUF_SIZE)
		nHead = uint32(0)
	)

	for {
		buflen := len(this_.rbuf)
		nHead = this_.nHead
		if (nHead == 0 && buflen < UINT32_SIZE) || (nHead > 0 && buflen < int(nHead)) {
			if this_.timeout > 0 {
				err := this_.tcpConn.SetReadDeadline(time.Now().Add(this_.timeout))
				if err != nil {
					return nil, err
				}
			}

			n, err := this_.tcpConn.Read(rbuf)
			if err != nil {
				return nil, err
			}

			this_.rbuf = append(this_.rbuf, rbuf[:n]...)
		}

		if nHead == 0 && len(this_.rbuf) >= UINT32_SIZE {
			nHead = binary.BigEndian.Uint32(this_.rbuf[:UINT32_SIZE])
			if nHead > MAX_BUF_SIZE || nHead == 0 {
				return nil, ErrInvalidBufSize
			}

			this_.rbuf = this_.rbuf[UINT32_SIZE:]
			this_.nHead = nHead
		}

		if nHead > 0 && len(this_.rbuf) >= int(nHead) {
			ret := this_.rbuf[:nHead]
			this_.rbuf = this_.rbuf[nHead:]
			this_.nHead = 0
			return ret, nil
		}
	}
}

func (this_ *Sess) wsRead() ([]byte, error) {
	if this_.wsConn == nil {
		log.Fatal("Sess.wsConn is nil")
	}

	_, data, err := this_.wsConn.ReadMessage()
	return data, err
}

func (this_ *Sess) tcpWrite(data []byte) (int, error) {
	return write(this_.tcpConn, data, this_.timeout)
}

func (this_ *Sess) wsWrite(data []byte) (int, error) {
	err := this_.wsConn.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		return -1, err
	}

	return len(data), nil
}
