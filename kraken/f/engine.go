package f

import (
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/nw"
	"github.com/iegad/gox/kraken/f/handlers"
	"github.com/iegad/gox/kraken/m"
	"github.com/iegad/gox/pb"
	"google.golang.org/protobuf/proto"
)

type engine struct {
	handlers map[int32]m.Handler
}

func newEngine() *engine {
	this_ := &engine{
		handlers: make(map[int32]m.Handler),
	}

	this_.addHandler(pb.MID_UserLoginReq, handlers.UserLogin)

	return this_
}

func (this_ *engine) addHandler(mid int32, h m.Handler) {
	if _, ok := this_.handlers[mid]; ok {
		log.Fatal("%v has already exists", mid)
	}

	this_.handlers[mid] = h
}

func (this_ *engine) Info() *nw.EngineInfo {
	return &nw.EngineInfo{Name: "Kraken.Front Service"}
}

func (this_ *engine) OnConnected(sess *nw.Sess) error {
	m.Players.AddSession(sess)
	log.Info("%v[%v] 连接成功", sess.Protocol(), sess.RemoteAddr().String())
	return nil
}

func (this_ *engine) OnDisconnected(sess *nw.Sess) {
	m.Players.RemoveSession(sess.RemoteAddr().String())
	log.Info("%v[%v] 连接断开", sess.Protocol(), sess.RemoteAddr().String())
}

func (this_ *engine) OnData(sess *nw.Sess, data []byte) error {
	pack := &pb.Package{}
	err := proto.Unmarshal(data, pack)
	if err != nil {
		return err
	}

	if pack.MessageID == 0 {
		return m.Err_F_MessageIDInvalid
	}

	if len(pack.NodeCode) == 0 {
		if handler, ok := this_.handlers[pack.MessageID]; ok {
			err = handler(sess, pack)
			return err
		} else {
			return m.Err_F_MessageIDInvalid
		}
	}

	return nil
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

	m.Players.Clear()
}
