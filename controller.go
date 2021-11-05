package main

import (
	"conjur/module"
	"conjur/server"
)

type Controller struct {
	*module.Module
	Server *server.Server
}

func MakeController() Controller {
	controller := Controller{}
	controller.Module = module.MakeModule("Controller", controller.Init, controller.DeInit)
	controller.Server = server.MakeServer()
	return controller
}

func (controller *Controller) Init() {
	controller.Server.Start()
}

func (controller *Controller) DeInit() {
	controller.Server.Stop()
}
