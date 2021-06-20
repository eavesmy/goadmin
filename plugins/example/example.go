package example

import (
	c "github.com/eavesmy/goadmin/modules/config"
	"github.com/eavesmy/goadmin/modules/service"
	"github.com/eavesmy/goadmin/plugins"
)

type Example struct {
	*plugins.Base
}

func NewExample() *Example {
	return &Example{
		Base: &plugins.Base{PlugName: "example"},
	}
}

func (e *Example) InitPlugin(srv service.List) {
	e.InitBase(srv, "example")
	e.App = e.initRouter(c.Prefix(), srv)
}
