package m

import (
	"sync"

	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/nw"
)

type NodeManager struct {
	nodeMap sync.Map
}

func (this_ *NodeManager) AddSession(sess *nw.Sess) {
	if sess == nil || sess.UserData != nil {
		log.Fatal("sess is invalid")
	}

	this_.nodeMap.Store(sess.RemoteAddr().String(), sess)
}

func (this_ *NodeManager) AddNode(nd *node) {
	if nd.NodeID <= 0 || len(nd.NodeCode) != 36 || nd.Sess == nil || nd.Sess.UserData == nil {
		log.Fatal("n is invalid")
	}

	this_.nodeMap.Delete(nd.Sess.RemoteAddr().String())
	this_.nodeMap.Store(nd.NodeCode, nd)
}

func (this_ *NodeManager) RemoveSession(remoteAddr string) {
	if len(remoteAddr) == 0 {
		log.Fatal("remoteAddr is invalid")
	}

	if v, ok := this_.nodeMap.LoadAndDelete(remoteAddr); ok {
		v.(*nw.Sess).Shutdown()
	}
}

func (this_ *NodeManager) RemovePlayer(nodeCode string) {
	if len(nodeCode) != 36 {
		log.Fatal("nodeCode is invalid")
	}

	if v, ok := this_.nodeMap.LoadAndDelete(nodeCode); ok {
		v.(*node).Sess.Shutdown()
	}
}

func (this_ *NodeManager) GetNodeByCode(nodeCode string) *node {
	if len(nodeCode) != 36 {
		log.Fatal("nodeCode is invalid")
	}

	if v, ok := this_.nodeMap.Load(nodeCode); ok {
		return v.(*node)
	}

	return nil
}

func (this_ *NodeManager) Clear() {
	this_.nodeMap.Range(func(key, value any) bool {
		k := key.(string)
		if len(k) == 36 {
			value.(*node).Sess.Shutdown()
		} else {
			value.(*nw.Sess).Shutdown()
		}
		return true
	})

	this_.nodeMap = sync.Map{}
}
