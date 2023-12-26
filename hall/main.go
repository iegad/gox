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
	go m.RunKeepAlived()
	m.InitHall()

	m.HallNode.Run(4)
}
