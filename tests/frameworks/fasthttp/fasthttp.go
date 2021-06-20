package fasthttp

import (
	// add fasthttp adapter
	ada "github.com/eavesmy/goadmin/adapter/fasthttp"
	// add mysql driver
	_ "github.com/eavesmy/goadmin/modules/db/drivers/mysql"
	// add postgresql driver
	_ "github.com/eavesmy/goadmin/modules/db/drivers/postgres"
	// add sqlite driver
	_ "github.com/eavesmy/goadmin/modules/db/drivers/sqlite"
	// add mssql driver
	_ "github.com/eavesmy/goadmin/modules/db/drivers/mssql"
	// add adminlte ui theme
	_ "github.com/GoAdminGroup/themes/adminlte"

	"os"

	"github.com/eavesmy/goadmin/engine"
	"github.com/eavesmy/goadmin/modules/config"
	"github.com/eavesmy/goadmin/modules/language"
	"github.com/eavesmy/goadmin/plugins/admin"
	"github.com/eavesmy/goadmin/plugins/admin/modules/table"
	"github.com/eavesmy/goadmin/template"
	"github.com/eavesmy/goadmin/template/chartjs"
	"github.com/eavesmy/goadmin/tests/tables"
	"github.com/GoAdminGroup/themes/adminlte"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func newHandler() fasthttp.RequestHandler {
	router := fasthttprouter.New()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators).AddDisplayFilterXssJsFilter()
	adminPlugin.AddGenerator("user", tables.GetUserTable)

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin).
		Use(router); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return func(ctx *fasthttp.RequestCtx) {
		router.Handler(ctx)
	}
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) fasthttp.RequestHandler {
	router := fasthttprouter.New()

	eng := engine.Default()

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfig(&config.Config{
		Databases: dbs,
		UrlPrefix: "admin",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:    language.EN,
		IndexUrl:    "/",
		Debug:       true,
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}).
		AddAdapter(new(ada.Fasthttp)).
		AddGenerators(gens).
		Use(router); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return func(ctx *fasthttp.RequestCtx) {
		router.Handler(ctx)
	}
}
