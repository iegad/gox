package f

import (
	"log"

	"github.com/iegad/gox/frm/nw"
)

var Service *service

type service struct {
	Engine *engine
	ios    *nw.IOService
}

func Init(cfg *nw.IOServiceConfig) error {
	if cfg == nil {
		log.Fatal("cfg is nil")
	}

	if Service != nil {
		log.Fatal("front's service is already initialized")
	}

	e := newEngine()

	ios, err := nw.NewIOService(cfg, e)
	if err != nil {
		return err
	}

	Service = &service{
		Engine: e,
		ios:    ios,
	}

	return nil
}

func (this_ *service) Shutdown() {
	this_.ios.Shutdown()
}

func (this_ *service) Run() error {
	return this_.ios.Run()
}

func (this_ *service) Info() *nw.IOServiceInfo {
	return this_.ios.Info()
}
