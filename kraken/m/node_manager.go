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
	var nodes *sync.Map
	v, ok := this_.nodeMap.Load(nd.NodeID)
	if !ok {
		nodes := &sync.Map{}
		nodes.Store(nd.NodeCode, nd)
		this_.nodeMap.Store(nd.NodeID, nodes)
	} else {
		nodes = v.(*sync.Map)
	}

	nodes.Delete(nd.NodeCode)
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

func (this_ *NodeManager) RemoveNode(nodeID int32, nodeCode string) {
	if len(nodeCode) != 36 {
		log.Fatal("nodeCode is invalid")
	}

	var nodes *sync.Map
	if v, ok := this_.nodeMap.Load(nodeID); ok {
		nodes = v.(*sync.Map)
		if v, ok := nodes.LoadAndDelete(nodeCode); ok {
			v.(*node).Sess.Shutdown()
		}
	}
}

func (this_ *NodeManager) GetNode(nodeID int32, nodeCode string) *node {
	if len(nodeCode) != 36 {
		log.Fatal("nodeCode is invalid")
	}

	if v, ok := this_.nodeMap.Load(nodeID); ok {
		if v, ok = v.(*sync.Map).Load(nodeCode); ok {
			return v.(*node)
		}
	}

	return nil
}

func (this_ *NodeManager) Clear() {
	this_.nodeMap.Range(func(key, value any) bool {
		k := key.(string)
		if len(k) == 36 {
			nodes := value.(*sync.Map)
			nodes.Range(func(key, value any) bool {
				value.(*node).Sess.Shutdown()
				return true
			})
		} else {
			value.(*nw.Sess).Shutdown()
		}
		return true
	})

	this_.nodeMap = sync.Map{}
}
