package f

import (
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/nw"
	"github.com/iegad/gox/kraken/m"
)

type engine struct {
	PlayerManager m.PlayerManager
}

func (this_ *engine) Info() *nw.EngineInfo {
	return &nw.EngineInfo{Name: "Kraken.Front Service"}
}

func (this_ *engine) OnConnected(sess *nw.Sess) error {
	this_.PlayerManager.AddSession(sess)
	log.Info("[%v] is connected", sess.RemoteAddr().String())
	return nil
}

func (this_ *engine) OnDisconnected(sess *nw.Sess) {
	this_.PlayerManager.RemoveSession(sess.RemoteAddr().String())
	log.Info("[%v] is disconnected", sess.RemoteAddr().String())
}

func (this_ *engine) OnData(sess *nw.Sess, data []byte) error {
	log.Info(string(data))
	_, err := sess.Write(data)
	return err
}

func (this_ *engine) OnRun(iosvc *nw.IOService) error {
	tcpAddr := iosvc.TcpAddr()
	if tcpAddr != nil {
		log.Info("front's service tcp[%v] is running ...", tcpAddr.String())
	}

	wsAddr := iosvc.WsAddr()
	if wsAddr != nil {
		log.Info("front's service ws[%v] is running ...", wsAddr.String())
	}

	return nil
}

func (this_ *engine) OnStopped(iosvc *nw.IOService) {
	tcpAddr := iosvc.TcpAddr()
	if tcpAddr != nil {
		log.Info("front's service tcp[%v] has stopped !!!", tcpAddr.String())
	}

	wsAddr := iosvc.WsAddr()
	if wsAddr != nil {
		log.Info("front's service ws[%v] has stopped !!!", wsAddr.String())
	}

	this_.PlayerManager.Clear()
}
