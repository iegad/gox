package main

import (
	"fmt"
	"sync"

	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/nw"
)

const nTimes = 100000

func main() {
	c, err := nw.NewTcpClient("", "127.0.0.1:6688", 0)
	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < nTimes; i++ {
			str := fmt.Sprintf("1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890: %d", i)
			_, err = c.TcpWrite([]byte(str))
			if err != nil {
				log.Error(err)
				break
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < nTimes; i++ {
			data, err := c.TcpRead()
			if err != nil {
				log.Error(err)
				break
			}
			log.Info(string(data))
		}
	}()

	wg.Wait()
}
