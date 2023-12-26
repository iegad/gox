package m

import (
	"github.com/iegad/gox/frm/proxy"
	"github.com/iegad/gox/hall/conf"
)

var HallNode *proxy.Node

func InitHall() {
	HallNode = proxy.NewNode(conf.Instance.NodeUID)
}
