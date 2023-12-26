package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/iegad/gox/hall/conf"
	"github.com/iegad/gox/hall/m"
)

func main() {
	var (
		sigs = make(chan os.Signal, 1)
	)

	signal.Notify(sigs, syscall.SIGINT)
	go func() {
		<-sigs
		os.Exit(0)
	}()

	conf.Init("config.yaml")
	m.InitHall()
	go m.RunKeepAlived()

	m.HallNode.Run(4)
}
