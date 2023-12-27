package m

import (
	"github.com/iegad/gox/frm/svc"
	"github.com/iegad/gox/hall/conf"
)

var HallNode *svc.Node

func InitHall() {
	HallNode = svc.NewNode(conf.Instance.NodeUID)
}
