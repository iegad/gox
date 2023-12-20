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
	log.Info("%v[%v] 连接成功", sess.Protocol(), sess.RemoteAddr().String())
	return nil
}

func (this_ *engine) OnDisconnected(sess *nw.Sess) {
	this_.PlayerManager.RemoveSession(sess.RemoteAddr().String())
	log.Info("%v[%v] 连接断开", sess.Protocol(), sess.RemoteAddr().String())
}

func (this_ *engine) OnData(sess *nw.Sess, data []byte) error {
	log.Info(string(data))
	_, err := sess.Write(data)
	return err
}

func (this_ *engine) OnRun(iosvc *nw.IOService) error {
	tcpAddr := iosvc.TcpAddr()
	if tcpAddr != nil {
		log.Info("Front Service TCP[%v] 服务开启 ...", tcpAddr.String())
	}

	wsAddr := iosvc.WsAddr()
	if wsAddr != nil {
		log.Info("Front Service ws[%v] 服务开启 ...", wsAddr.String())
	}

	return nil
}

func (this_ *engine) OnStopped(iosvc *nw.IOService) {
	tcpAddr := iosvc.TcpAddr()
	if tcpAddr != nil {
		log.Info("Front Service TCP[%v] 服务关闭", tcpAddr.String())
	}

	wsAddr := iosvc.WsAddr()
	if wsAddr != nil {
		log.Info("Front Service ws[%v] 服务关闭", wsAddr.String())
	}

	this_.PlayerManager.Clear()
}
