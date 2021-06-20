package beego

import (
	// add beego adapter
	_ "github.com/eavesmy/goadmin/adapter/beego"
	"github.com/eavesmy/goadmin/modules/config"
	"github.com/eavesmy/goadmin/modules/language"
	"github.com/eavesmy/goadmin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/themes/adminlte"

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

	"net/http"
	"os"

	"github.com/eavesmy/goadmin/engine"
	"github.com/eavesmy/goadmin/plugins/admin"
	"github.com/eavesmy/goadmin/plugins/example"
	"github.com/eavesmy/goadmin/template"
	"github.com/eavesmy/goadmin/template/chartjs"
	"github.com/eavesmy/goadmin/tests/tables"
	"github.com/astaxie/beego"
)

func newHandler() http.Handler {

	app := beego.NewApp()

	eng := engine.Default()
	adminPlugin := admin.NewAdmin(tables.Generators)
	adminPlugin.AddGenerator("user", tables.GetUserTable)

	examplePlugin := example.NewExample()

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(app); err != nil {
		panic(err)
	}

	template.AddComp(chartjs.NewChart())

	eng.HTML("GET", "/admin", tables.GetContent)

	beego.BConfig.Listen.HTTPAddr = "127.0.0.1"
	beego.BConfig.Listen.HTTPPort = 9087

	return app.Handlers
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {

	app := beego.NewApp()

	eng := engine.Default()
	adminPlugin := admin.NewAdmin(gens)

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
		AddPlugins(adminPlugin).Use(app); err != nil {
		panic(err)
	}

	template.AddComp(chartjs.NewChart())

	eng.HTML("GET", "/admin", tables.GetContent)

	beego.BConfig.Listen.HTTPAddr = "127.0.0.1"
	beego.BConfig.Listen.HTTPPort = 9087

	return app.Handlers
}
