package main

import (
	"github.com/google/uuid"
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/node"
)

func main() {
	uid := uuid.New()
	nd := node.NewNode(1, uid[:])
	err := nd.AddProxy("127.0.0.1:6688")
	if err != nil {
		log.Fatal(err)
	}

	nd.Run(4)
}
