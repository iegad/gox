package ios

import (
	"sync"

	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/nw"
)

var Front *nw.IOService

type front struct {
	sessMap sync.Map
}

func (this_ *front) Info() *nw.EngineInfo {
	return &nw.EngineInfo{Name: "EchoEngine"}
}

func (this_ *front) OnConnected(sess *nw.Sess) error {
	this_.sessMap.Store(sess.RemoteAddr().String(), sess)
	log.Info("[%v] is connected", sess.RemoteAddr().String())
	return nil
}

func (this_ *front) OnDisconnected(sess *nw.Sess) {
	this_.sessMap.Delete(sess.RemoteAddr().String())
	log.Info("[%v] is disconnected", sess.RemoteAddr().String())
}

func (this_ *front) OnData(sess *nw.Sess, data []byte) error {
	log.Info(string(data))
	_, err := sess.Write(data)
	return err
}

func (this_ *front) OnRun(iosvc *nw.IOService) error {
	log.Info("-------------- server run ----------------")

	log.Info(iosvc.Info())
	return nil
}

func (this_ *front) OnStopped(iosvc *nw.IOService) {
	log.Info("-------------- server stopped ----------------")
	log.Info(iosvc.Info())
}

func (this_ *front) OnShutdown(iosvc *nw.IOService) {
	this_.sessMap.Range(func(key, value any) bool {
		value.(*nw.Sess).Shutdown()
		return true
	})
}

func InitFront(cfg *nw.IOServiceConfig) error {
	if cfg == nil {
		log.Fatal("cfg is nil")
	}

	var err error

	Front, err = nw.NewIOService(cfg, &front{})
	if err != nil {
		return err
	}

	return nil
}
