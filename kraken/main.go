package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/kraken/b"
	"github.com/iegad/gox/kraken/conf"
	"github.com/iegad/gox/kraken/f"
	"github.com/iegad/gox/kraken/m"
	"github.com/iegad/gox/kraken/manager"
)

func main() {
	m.Init()

	conf.LoadConfig("config.yml")
	log.Info("config.yml 配置加载成功")

	err := f.Init(conf.Instance.Front)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Front Service 初始化完成")

	err = b.Init(conf.Instance.Backend)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Backend Service 初始化完成")

	err = manager.Init(conf.Instance.ManangerHost)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Manager 初始化完成")

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

	log.Info("Kraken [%v] 服务开启 ...", conf.Instance.ManangerHost)
	manager.Instance.Run()
	<-done
	log.Info("Kraken [%v] 服务关闭", conf.Instance.ManangerHost)
}
