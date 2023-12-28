package main

import (
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/nw"
)

type EchoEngine struct {
}

func (this_ *EchoEngine) Info() *nw.EngineInfo {
	return &nw.EngineInfo{
		Name: "EchoEngine",
	}
}

func (this_ *EchoEngine) OnConnected(sess *nw.Sess) error {
	log.Info("%v has connected", sess.RemoteAddr().String())
	return nil
}

func (this_ *EchoEngine) OnDisconnected(sess *nw.Sess) {
	log.Info("%v has disconnected", sess.RemoteAddr().String())
}

func (this_ *EchoEngine) OnData(sess *nw.Sess, data []byte) error {
	log.Info(string(data))
	_, err := sess.Write(data)
	return err
}

func (this_ *EchoEngine) OnRun(iosvc *nw.IOService) error {
	return nil
}

func (this_ *EchoEngine) OnStopped(iosvc *nw.IOService) {

}

func main() {
	ios, err := nw.NewIOService(&nw.IOServiceConfig{
		TcpEndpoint: ":8080",
		Timeout:     10,
	}, &EchoEngine{})
	if err != nil {
		log.Fatal(err)
	}

	ios.RunLoop()
}
