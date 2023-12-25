package nw

import (
	"encoding/binary"
	"net"
	"time"

	"github.com/iegad/gox/frm/log"
)

type Client struct {
	tcpConn  *net.TCPConn
	timeout  time.Duration
	nHead    uint32
	rbuf     []byte
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
		nHead:   0,
		rbuf:    []byte{},
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
			nHead = binary.BigEndian.Uint32(this_.rbuf[:UINT32_SIZE]) ^ _HeaderKey
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
