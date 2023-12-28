package nw

import (
	"bufio"
	"encoding/binary"
	"net"
	"time"

	"github.com/iegad/gox/frm/log"
)

type Client struct {
	tcpConn  *net.TCPConn
	timeout  time.Duration
	header   uint32
	reader   *bufio.Reader
	UserData interface{}
}

func NewTcpClient(lep, rep string, timeout time.Duration) (*Client, error) {
	var (
		conn  *net.TCPConn
		laddr *net.TCPAddr
		raddr *net.TCPAddr
		err   error
	)

	raddr, err = net.ResolveTCPAddr("tcp", rep)
	if err != nil {
		return nil, err
	}

	if len(lep) > 0 {
		laddr, err = net.ResolveTCPAddr("tcp", lep)
		if err != nil {
			return nil, err
		}
	}

	conn, err = net.DialTCP("tcp", laddr, raddr)
	if err != nil {
		return nil, err
	}

	return &Client{
		tcpConn: conn,
		timeout: timeout,
		reader:  bufio.NewReader(conn),
	}, nil
}

func (this_ *Client) TcpWrite(data []byte) (int, error) {
	if this_.tcpConn == nil {
		log.Fatal("Client.tcpConn is nil")
	}

	return write(this_.tcpConn, data, this_.timeout)
}

func (this_ *Client) Raw() net.Conn {
	return this_.tcpConn
}

func (this_ *Client) TcpRead() ([]byte, error) {
	if this_.tcpConn == nil {
		log.Fatal("Client.tcpConn is nil")
	}

	if this_.timeout > 0 {
		err := this_.tcpConn.SetReadDeadline(time.Now().Add(this_.timeout))
		if err != nil {
			return nil, err
		}
	}

	header := this_.header
	if header == 0 {
		peek, err := this_.reader.Peek(UINT32_SIZE)
		if err != nil {
			return nil, err
		}

		header = binary.BigEndian.Uint32(peek) ^ __HEADER_KEY_
		if header == 0 || header > MAX_BUF_SIZE {
			return nil, ErrInvalidBufSize
		}

		this_.header = header
	}

	buflen := header + uint32(UINT32_SIZE)
	if this_.reader.Buffered() < int(buflen) {
		return nil, nil
	}

	rbuf := make([]byte, buflen)
	_, err := this_.reader.Read(rbuf)
	if err != nil {
		return nil, err
	}

	this_.header = 0
	return rbuf[UINT32_SIZE:], nil
}
