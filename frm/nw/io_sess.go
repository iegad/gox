package nw

import (
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

func (this_ *Sess) Protocol() string {
	if this_.tcpConn == nil {
		return "ws"
	}

	return "tcp"
}

func (this_ *Sess) Shutdown() {
	if this_.tcpConn != nil {
		this_.tcpConn.Close()
	}
}

func (this_ *Sess) Write(data []byte) (int, error) {
	if this_.tcpConn != nil {
		return this_.TcpWrite(data)
	}

	return this_.WsWrite(data)
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

func (this_ *Sess) TcpWrite(data []byte) (int, error) {
	return write(this_.tcpConn, data, this_.timeout)
}

func (this_ *Sess) WsWrite(data []byte) (int, error) {
	err := this_.wsConn.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		return -1, err
	}

	return len(data), nil
}
