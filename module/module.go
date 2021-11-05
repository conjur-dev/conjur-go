package module

import (
	"conjur/logger"
	"fmt"
)

type Module struct {
	Name   string
	Init   func()
	DeInit func()
}

func MakeModule(name string, starter func(), stopper func()) *Module {
	return &Module{
		Name:   name,
		Init:   starter,
		DeInit: stopper,
	}
}

func (module *Module) Start() {
	logger.Info(fmt.Sprintf("%s::Start() ...", module.Name))
	module.Init()
	logger.Info(fmt.Sprintf("%s::Start() = 0", module.Name))
}

func (module *Module) Stop() {
	logger.Info(fmt.Sprintf("%s::Stop() ...", module.Name))
	module.DeInit()
	logger.Info(fmt.Sprintf("%s::Stop() = 0", module.Name))
}
