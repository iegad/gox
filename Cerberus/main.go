package main

import (
	"github.com/iegad/gox/cerberus/conf"
	"github.com/iegad/gox/cerberus/m"
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/web"
)

func main() {
	conf.Init("config.yml")
	go m.RunGetAllNodes()

	server, err := web.NewServer(conf.Instance.Host, true)
	if err != nil {
		log.Fatal(err)
	}

	log.Error(server.Run())
}
