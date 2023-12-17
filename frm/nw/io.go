package nw

import (
	"encoding/binary"
	"net"
	"time"

	"github.com/iegad/gox/frm/log"
)

func write(conn net.Conn, data []byte, timeout time.Duration) (int, error) {

	if conn == nil {
		log.Fatal("conn is nil")
	}

	if timeout > 0 {
		err := conn.SetWriteDeadline(time.Now().Add(timeout))
		if err != nil {
			return -1, err
		}
	}

	dlen := len(data)
	wbuf := make([]byte, dlen+UINT32_SIZE)
	binary.BigEndian.PutUint32(wbuf[:UINT32_SIZE], uint32(dlen))
	copy(wbuf[UINT32_SIZE:], data)
	return conn.Write(wbuf)
}
