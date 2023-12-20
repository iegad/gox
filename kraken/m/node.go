package m

import "github.com/iegad/gox/frm/nw"

type node struct {
	NodeID   int32
	NodeCode string
	Sess     *nw.Sess
}

func NewNode(nodeID int32, nodeCode string, sess *nw.Sess) *node {
	return &node{
		NodeID:   nodeID,
		NodeCode: nodeCode,
		Sess:     sess,
	}
}
