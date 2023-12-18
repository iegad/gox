package ios

import (
	"sync"

	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/nw"
)

var Backend *nw.IOService

type backend struct {
	sessMap sync.Map
}

func (this_ *backend) Info() *nw.EngineInfo {
	return &nw.EngineInfo{Name: "EchoEngine"}
}

func (this_ *backend) OnConnected(sess *nw.Sess) error {
	this_.sessMap.Store(sess.RemoteAddr().String(), sess)
	log.Info("[%v] is connected", sess.RemoteAddr().String())
	return nil
}

func (this_ *backend) OnDisconnected(sess *nw.Sess) {
	this_.sessMap.Delete(sess.RemoteAddr().String())
	log.Info("[%v] is disconnected", sess.RemoteAddr().String())
}

func (this_ *backend) OnData(sess *nw.Sess, data []byte) error {
	log.Info(string(data))
	_, err := sess.Write(data)
	return err
}

func (this_ *backend) OnRun(iosvc *nw.IOService) error {
	tcpAddr := iosvc.TcpAddr()
	if tcpAddr != nil {
		log.Info("backend's service tcp[%v] is running ...", tcpAddr.String())
	}

	return nil
}

func (this_ *backend) OnStopped(iosvc *nw.IOService) {
	tcpAddr := iosvc.TcpAddr()
	if tcpAddr != nil {
		log.Info("backend's service tcp[%v] has stopped !!!", tcpAddr.String())
	}
}

func (this_ *backend) OnShutdown(iosvc *nw.IOService) {
	this_.sessMap.Range(func(key, value any) bool {
		value.(*nw.Sess).Shutdown()
		return true
	})
}

func InitBackend(cfg *nw.IOServiceConfig) error {
	if cfg == nil {
		log.Fatal("cfg is nil")
	}

	var err error
	Backend, err = nw.NewIOService(cfg, &backend{})
	return err
}
