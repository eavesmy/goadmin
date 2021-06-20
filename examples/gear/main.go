package main

import (
	"log"
	"os"
	"os/signal"

<<<<<<< HEAD
	_ "github.com/GoAdminGroup/go-admin/adapter/gear"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
	_ "github.com/GoAdminGroup/themes/sword"
	"github.com/eavesmy/gear"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/example"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/themes/adminlte"
=======
	_ "github.com/GoAdminGroup/themes/sword"
	_ "github.com/eavesmy/goadmin/adapter/gear"
	_ "github.com/eavesmy/goadmin/modules/db/drivers/postgres"

	"github.com/GoAdminGroup/themes/adminlte"
	"github.com/eavesmy/gear"
	"github.com/eavesmy/gear/middleware/static"
	"github.com/eavesmy/goadmin/engine"
	"github.com/eavesmy/goadmin/examples/datamodel"
	"github.com/eavesmy/goadmin/modules/config"
	"github.com/eavesmy/goadmin/modules/language"
	"github.com/eavesmy/goadmin/plugins/example"
	"github.com/eavesmy/goadmin/template"
	"github.com/eavesmy/goadmin/template/chartjs"
>>>>>>> 2fe0a5673b0868e2c22ef57fcc1e5c8a07ee24e7
)

func main() {

	app := gear.New()

	e := engine.Default()

	cfg := config.Config{
		Env: config.EnvLocal,
		Databases: config.DatabaseList{
			"default": {
				Host:       "10.40.126.223",
				Port:       "5432",
				User:       "postgres",
				Pwd:        "23216340pgsqlDL",
				Name:       "godmin",
<<<<<<< HEAD
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     config.DriverPostgresql,
				//File:   "../datamodel/admin.db",
			},
		},
		UrlPrefix: "/",
=======
				MaxIdleCon: 5,
				MaxOpenCon: 10,
				Driver:     config.DriverPostgresql,

				//Driver: config.DriverSqlite,
				//File:   "../datamodel/admin.db",
			},
		},
>>>>>>> 2fe0a5673b0868e2c22ef57fcc1e5c8a07ee24e7
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
<<<<<<< HEAD
=======
		UrlPrefix:          "admin",
>>>>>>> 2fe0a5673b0868e2c22ef57fcc1e5c8a07ee24e7
		Language:           language.CN,
		IndexUrl:           "/",
		Debug:              true,
		AccessAssetsLogOff: true,
		Animation: config.PageAnimation{
			Type: "fadeInUp",
		},
		ColorScheme:       adminlte.ColorschemeSkinBlack,
		BootstrapFilePath: "./../datamodel/bootstrap.go",
	}

	template.AddComp(chartjs.NewChart())

	// customize a plugin

	examplePlugin := example.NewExample()

	// load from golang.Plugin
	//
	// examplePlugin := plugins.LoadFromPlugin("../datamodel/example.so")

	// customize the login page
	// example: https://github.com/GoAdminGroup/demo.go-admin.cn/blob/master/main.go#L39
	//
	// template.AddComp("login", datamodel.LoginPage)

	// load config from json file
	//
	// e.AddConfigFromJSON("../datamodel/config.json")

<<<<<<< HEAD
=======
	app.Use(static.New(static.Options{Root: "./uploads", Prefix: "uploads"}))

>>>>>>> 2fe0a5673b0868e2c22ef57fcc1e5c8a07ee24e7
	if err := e.AddConfig(&cfg).
		AddGenerators(datamodel.Generators).
		// add generator, first parameter is the url prefix of table when visit.
		// example:
		//
		// "user" => http://localhost:9033/admin/info/user
		//
		AddGenerator("user", datamodel.GetUserTable).
		AddDisplayFilterXssJsFilter().
		AddPlugins(examplePlugin).
		Use(app); err != nil {
		panic(err)
	}

	// customize your pages

	e.HTML("GET", "/admin", datamodel.GetContent)

	go func() {
<<<<<<< HEAD
		_ = app.Start(":20000")
=======
		app.Start(":8099")
>>>>>>> 2fe0a5673b0868e2c22ef57fcc1e5c8a07ee24e7
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
<<<<<<< HEAD
	e.PostgresqlConnection().Close()
=======
	e.MysqlConnection().Close()
>>>>>>> 2fe0a5673b0868e2c22ef57fcc1e5c8a07ee24e7
}
