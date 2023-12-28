package m

import (
	"sync"

	"github.com/iegad/gox/frm/biz"
	"github.com/iegad/gox/frm/log"
)

type nodeManager struct {
	proxyMap map[int32]*biz.ProxyInfo
	hallMap  map[int32]*biz.HallInfo
	mtx      sync.RWMutex
}

var NodeManager *nodeManager = &nodeManager{}

func (this_ *nodeManager) SetNodes(proxyList []*biz.ProxyInfo, hallList []*biz.HallInfo) {
	proxyMap := make(map[int32]*biz.ProxyInfo)
	hallMap := make(map[int32]*biz.HallInfo)

	if len(proxyList) > 0 {
		for _, item := range proxyList {
			log.Info("PROXY [%v:%v]", item.ChannelID, item.NodeCode)
			proxyMap[item.ChannelID] = item
		}
	}

	if len(hallList) > 0 {
		for _, item := range hallList {
			log.Info("HALL [%v:%v]", item.ChannelID, item.NodeCode)
			hallMap[item.ChannelID] = item
		}
	}

	this_.mtx.Lock()
	this_.proxyMap = proxyMap
	this_.hallMap = hallMap
	this_.mtx.Unlock()
}

func (this_ *nodeManager) GetHall(channelID int32) *biz.HallInfo {
	this_.mtx.RLock()
	defer this_.mtx.RUnlock()

	if v, ok := this_.hallMap[channelID]; ok {
		return v
	}

	return nil
}

func (this_ *nodeManager) GetProxy(channelID int32) *biz.ProxyInfo {
	this_.mtx.RLock()
	defer this_.mtx.RUnlock()

	if v, ok := this_.proxyMap[channelID]; ok {
		return v
	}

	return nil
}
