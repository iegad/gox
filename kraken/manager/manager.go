package manager

import (
	"sync"

	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/web"
	"github.com/iegad/gox/kraken/ios"
	"github.com/iegad/gox/kraken/manager/cgi"
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
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := ios.Front.Run()
		if err != nil {
			log.Error("front service: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := ios.Backend.Run()
		if err != nil {
			log.Error("backend service: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := this_.webServer.Run()
		if err != nil {
			log.Error("manager web: %v", err)
		}
	}()

	wg.Wait()
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
	Instance = tmp
	return nil
}
