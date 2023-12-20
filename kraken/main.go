package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/kraken/b"
	"github.com/iegad/gox/kraken/conf"
	"github.com/iegad/gox/kraken/f"
	"github.com/iegad/gox/kraken/manager"
)

func main() {
	err := conf.LoadConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = f.Init(conf.Instance.Front)
	if err != nil {
		log.Fatal(err)
	}

	err = b.Init(conf.Instance.Backend)
	if err != nil {
		log.Fatal(err)
	}

	err = manager.Init(conf.Instance.ManangerHost)
	if err != nil {
		log.Fatal(err)
	}

	var (
		sigs = make(chan os.Signal, 1)
		done = make(chan bool, 1)
	)

	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		<-sigs
		manager.Instance.Shutdown()
		done <- true
	}()

	manager.Instance.Run()
	<-done
	log.Info("Kraken is exit !!!")
}
