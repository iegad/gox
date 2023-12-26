package b

import (
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/nw"
	"github.com/iegad/gox/kraken/b/handlers"
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

	this_.addHandler(pb.MID_B_RegistNodeReq, handlers.RegistNode)

	return this_
}

func (this_ *engine) addHandler(mid int32, h m.Handler) {
	if _, ok := this_.handlers[mid]; ok {
		log.Fatal("%v has already exists", mid)
	}

	this_.handlers[mid] = h
}

func (this_ *engine) Info() *nw.EngineInfo {
	return &nw.EngineInfo{Name: "Kraken.Backend"}
}

func (this_ *engine) OnConnected(sess *nw.Sess) error {
	m.Nodes.AddSession(sess)
	return nil
}

func (this_ *engine) OnDisconnected(sess *nw.Sess) {
	m.Nodes.Remove(sess)
}

func (this_ *engine) OnData(sess *nw.Sess, data []byte) error {
	pack := &pb.Package{}
	err := proto.Unmarshal(data, pack)
	if err != nil {
		return err
	}

	if pack.MessageID == 0 {
		return m.Err_B_MessageIDInvalid
	}

	if len(pack.NodeCode) != 16 {
		return m.Err_B_NodeCodeInvalid
	}

	if handler, ok := this_.handlers[pack.MessageID]; ok {
		err = handler(sess, pack)
		return err
	} else {
		return m.Err_F_MessageIDInvalid
	}
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
