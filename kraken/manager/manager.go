package manager

import (
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/web"
	"github.com/iegad/gox/kraken/b"
	"github.com/iegad/gox/kraken/manager/cgi"
	"github.com/iegad/gox/kraken/manager/cgi/backend"
	"github.com/iegad/gox/kraken/manager/cgi/front"

	"github.com/iegad/gox/kraken/f"
)

var Instance *manager

type manager struct {
	webServer *web.Server
}

func (this_ *manager) Shutdown() {
	if f.Service != nil {
		f.Service.Shutdown()
	}

	if b.Service != nil {
		b.Service.Shutdown()
	}

	if this_.webServer != nil {
		this_.webServer.Shutdown()
	}
}

func (this_ *manager) Run() {
	go func() {
		err := f.Service.Run()
		if err != nil {
			log.Error("front service: %v", err)
		}
	}()

	go func() {
		err := b.Service.Run()
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
	tmp.webServer.Router().POST("/front/kick_session", front.KickSession)

	tmp.webServer.Router().POST("/backend/shutdown", backend.Shutdown)
	tmp.webServer.Router().POST("/backend/run", backend.Run)
	tmp.webServer.Router().POST("/backend/info", backend.Info)
	Instance = tmp
	return nil
}
