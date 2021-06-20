package gorilla

import (
	// add gorilla adapter
	_ "github.com/eavesmy/goadmin/adapter/gorilla"
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
	"github.com/gorilla/mux"
)

func newHandler() http.Handler {
	app := mux.NewRouter()
	eng := engine.Default()

	examplePlugin := example.NewExample()
	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(admin.NewAdmin(tables.Generators).
			AddGenerator("user", tables.GetUserTable), examplePlugin).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return app
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {
	app := mux.NewRouter()
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
		AddPlugins(admin.NewAdmin(gens)).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return app
}
