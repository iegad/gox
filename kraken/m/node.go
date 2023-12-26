package m

import (
	"sync"

	"github.com/google/uuid"
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/nw"
	"github.com/iegad/gox/pb"
)

type Handler func(sess *nw.Sess, in *pb.Package) error

type Node struct {
	NodeCode   string
	Idempotent int64
	Sess       *nw.Sess
}

func NewNode(nodeCode []byte, sess *nw.Sess) *Node {
	if len(nodeCode) != 16 {
		log.Fatal("nodeCode is invalid")
	}

	this_ := &Node{
		NodeCode:   uuid.UUID(nodeCode).String(),
		Idempotent: 0,
		Sess:       sess,
	}

	sess.UserData = this_
	return this_
}

type NodeManager struct {
	nodeMap sync.Map
}

func (this_ *NodeManager) AddSession(sess *nw.Sess) {
	if sess == nil || sess.UserData != nil {
		log.Fatal("sess is invalid")
	}

	this_.nodeMap.Store(sess.RemoteAddr().String(), sess)
}

func (this_ *NodeManager) AddNode(nd *Node) {
	log.Info(nd.NodeCode)
	if len(nd.NodeCode) != 36 || nd.Sess == nil || nd.Sess.UserData == nil {
		log.Fatal("n is invalid")
	}

	this_.nodeMap.Delete(nd.Sess.RemoteAddr().String())
	this_.nodeMap.Store(nd.NodeCode, nd)
	log.Info("node: %v 注册成功", nd.NodeCode)
}

func (this_ *NodeManager) Remove(sess *nw.Sess) {
	if sess.UserData == nil {
		if v, ok := this_.nodeMap.LoadAndDelete(sess.RemoteAddr().String()); ok {
			v.(*nw.Sess).Shutdown()
		}
	} else {
		nd := sess.UserData.(*Node)
		if v, ok := this_.nodeMap.LoadAndDelete(nd.NodeCode); ok {
			s := v.(*Node).Sess
			log.Info("Node [%v]: %v 已断开连接", s.RemoteAddr().String(), nd.NodeCode)
			s.Shutdown()
		}
	}
}

func (this_ *NodeManager) GetNode(nodeCode string) *Node {
	if len(nodeCode) != 36 {
		log.Fatal("nodeCode is invalid")
	}

	if v, ok := this_.nodeMap.Load(nodeCode); ok {
		return v.(*Node)
	}

	return nil
}

func (this_ *NodeManager) Clear() {
	this_.nodeMap.Range(func(key, value any) bool {
		k := key.(string)
		if len(k) == 36 {
			value.(*Node).Sess.Shutdown()
		} else {
			value.(*nw.Sess).Shutdown()
		}
		return true
	})

	this_.nodeMap = sync.Map{}
}
