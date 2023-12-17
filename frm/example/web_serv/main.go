package main

import (
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/nw"
	"github.com/iegad/gox/frm/web"
)

type EchoEngine struct {
	sessMap sync.Map
}

func (this_ *EchoEngine) Info() *nw.EngineInfo {
	return &nw.EngineInfo{Name: "EchoEngine"}
}

func (this_ *EchoEngine) OnConnected(sess *nw.Sess) error {
	this_.sessMap.Store(sess.RemoteAddr().String(), sess)
	log.Info("[%v] is connected", sess.RemoteAddr().String())
	return nil
}

func (this_ *EchoEngine) OnDisconnected(sess *nw.Sess) {
	this_.sessMap.Delete(sess.RemoteAddr().String())
	log.Info("[%v] is disconnected", sess.RemoteAddr().String())
}

func (this_ *EchoEngine) OnData(sess *nw.Sess, data []byte) error {
	log.Info(string(data))
	_, err := sess.Write(data)
	return err
}

func (this_ *EchoEngine) OnRun(iosvc *nw.IOService) error {
	log.Info("-------------- server run ----------------")

	log.Info(iosvc.Info())
	return nil
}

func (this_ *EchoEngine) OnStopped(iosvc *nw.IOService) {
	log.Info("-------------- server stopped ----------------")
	log.Info(iosvc.Info())
}

func (this_ *EchoEngine) OnShutdown(iosvc *nw.IOService) {
	this_.sessMap.Range(func(key, value any) bool {
		value.(*nw.Sess).Shutdown()
		return true
	})
}

func main() {
	server, err := web.NewServer(":8888", true)
	if err != nil {
		log.Error(err)
		return
	}

	iosvc, err := nw.DefaultIOService(&EchoEngine{})
	if err != nil {
		log.Error(err)
		return
	}

	server.Router().POST("/io_service/run", func(ctx *gin.Context) { go iosvc.Run() })
	server.Router().POST("/io_service/shutdown", func(ctx *gin.Context) { iosvc.Shutdown() })
	server.Router().GET("/io_service/info", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, iosvc.Info()) })

	var (
		sigs = make(chan os.Signal, 1)
		done = make(chan bool, 1)
	)

	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		<-sigs
		server.Shutdown()
		iosvc.Shutdown()
		log.Info("服务已关闭")
		done <- true
	}()

	log.Error(server.Run())
	<-done
}
