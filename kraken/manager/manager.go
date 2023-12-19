package manager

import (
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/web"
	"github.com/iegad/gox/kraken/ios"
	"github.com/iegad/gox/kraken/manager/cgi"
	"github.com/iegad/gox/kraken/manager/cgi/backend"
	"github.com/iegad/gox/kraken/manager/cgi/front"
)

var Instance *manager

type manager struct {
	webServer *web.Server
}

func (this_ *manager) Shutdown() {
	if ios.Front != nil {
		ios.Front.Shutdown()
	}

	if ios.Backend != nil {
		ios.Backend.Shutdown()
	}

	if this_.webServer != nil {
		this_.webServer.Shutdown()
	}
}

func (this_ *manager) Run() {
	go func() {
		err := ios.Front.Run()
		if err != nil {
			log.Error("front service: %v", err)
		}
	}()

	go func() {
		err := ios.Backend.Run()
		if err != nil {
			log.Error("backend service: %v", err)
		}
	}()

	err := this_.webServer.Run()
	if err != nil {
		log.Error("manager web: %v", err)
	}
}

func Init(host string) error {
	var (
		err error
		tmp = &manager{
			webServer: nil,
		}
	)

	tmp.webServer, err = web.NewServer(host, true)
	if err != nil {
		return err
	}

	tmp.webServer.Router().POST("/login", cgi.Login)
	tmp.webServer.Router().POST("/front/shutdown", front.Shutdown)
	tmp.webServer.Router().POST("/front/run", front.Run)
	tmp.webServer.Router().POST("/front/info", front.Info)
	tmp.webServer.Router().POST("/backend/shutdown", backend.Shutdown)
	tmp.webServer.Router().POST("/backend/run", backend.Run)
	tmp.webServer.Router().POST("/backend/info", backend.Info)
	Instance = tmp
	return nil
}
