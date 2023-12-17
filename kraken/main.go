package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/web"
	"github.com/iegad/gox/kraken/conf"
	"github.com/iegad/gox/kraken/ios"
)

func main() {
	err := conf.LoadConfig("./config.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = ios.InitFront(conf.Instance.Front)
	if err != nil {
		log.Fatal(err)
	}

	server, err := web.NewServer(":8888", true)
	if err != nil {
		log.Error(err)
		return
	}

	// server.Router().POST("/io_service/run", func(ctx *gin.Context) { go iosvc.Run() })
	// server.Router().POST("/io_service/shutdown", func(ctx *gin.Context) { iosvc.Shutdown() })
	// server.Router().GET("/io_service/info", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, iosvc.Info()) })

	var (
		sigs = make(chan os.Signal, 1)
		done = make(chan bool, 1)
	)

	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		<-sigs
		server.Shutdown()
		// iosvc.Shutdown()
		log.Info("服务已关闭")
		done <- true
	}()

	log.Error(server.Run())
	<-done
}
