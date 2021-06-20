package example

import (
	"github.com/eavesmy/goadmin/context"
	"github.com/eavesmy/goadmin/modules/auth"
	"github.com/eavesmy/goadmin/modules/db"
	"github.com/eavesmy/goadmin/modules/service"
)

func (e *Example) initRouter(prefix string, srv service.List) *context.App {

	app := context.NewApp()
	route := app.Group(prefix)
	route.GET("/example", auth.Middleware(db.GetConnection(srv)), e.TestHandler)

	return app
}
