package b

import (
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/nw"
	"github.com/iegad/gox/kraken/m"
)

type engine struct {
}

func (this_ *engine) Info() *nw.EngineInfo {
	return &nw.EngineInfo{Name: "Kraken.Backend"}
}

func (this_ *engine) OnConnected(sess *nw.Sess) error {
	m.Nodes.AddSession(sess)
	log.Info("[%v] 连接成功", sess.RemoteAddr().String())
	return nil
}

func (this_ *engine) OnDisconnected(sess *nw.Sess) {
	m.Nodes.RemoveSession(sess.RemoteAddr().String())
	log.Info("[%v] 断开连接", sess.RemoteAddr().String())
}

func (this_ *engine) OnData(sess *nw.Sess, data []byte) error {
	log.Info(string(data))
	_, err := sess.Write(data)
	return err
}

func (this_ *engine) OnRun(iosvc *nw.IOService) error {
	tcpAddr := iosvc.TcpAddr()
	if tcpAddr != nil {
		log.Info("Backend Service TCP[%v] 服务开启 ...", tcpAddr.String())
	}

	return nil
}

func (this_ *engine) OnStopped(iosvc *nw.IOService) {
	tcpAddr := iosvc.TcpAddr()
	if tcpAddr != nil {
		log.Info("Backend Service TCP[%v] 服务关闭", tcpAddr.String())
	}

	m.Nodes.Clear()
}
