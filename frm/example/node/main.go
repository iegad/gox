package main

import (
	"github.com/google/uuid"
	"github.com/iegad/gox/frm/proxy"
)

func main() {
	uid := uuid.New()
	nd := proxy.NewNode(uid[:])
	nd.AddProxy("ae906ff9-552b-4f63-b807-95a1195deddf", "127.0.0.1:6688")
	nd.Run(4)
}
