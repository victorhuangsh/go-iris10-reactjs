package main

import (
	"./api/user"
	"./modules"
	"./routes"
	"./const_vars"
	"./appconfigs"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/mvc"
	"net/http"
	"time"
	"os"

)

//var app *iris.Framework

var app *iris.Application
var cfg_all map[string]interface{}
var appcfg *appconfigs.AppConfigs

func init(){

}

func main() {

	app = iris.New()
	//cfg_all = initConfigration(app)

	appcfg = appconfigs.GetInstance(app)

	// set the router we want to use
	//app.Adapt(httprouter.New())

	/*
	// adapt a new logger which will print the dev messages(mostly errors)
	// and panic on production messages (by-default only fatal errors are printed via ProdMode)
	app.Adapt(iris.LoggerPolicy(func(mode iris.LogMode, msg string) {
		if mode == iris.DevMode {
			log.Printf(msg)
		} else if mode == iris.ProdMode {
			panic(msg)
		}
	})) // or use app.Adapt(iris.DevLogger()) to print only DevMode messages to the os.Stdout

	*/
	app.Logger().SetOutput(newLogFile( ) )
	//app.Logger().SetLevel("debug")

	// set the favicon
	app.Favicon("./frontend/public/images/favicon.ico")

	// configure public folders
	modules.Public(app)

	// set the global middlewares
	app.Use(logger.New())

	// set the custom errors
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.ViewData("Title", http.StatusText(iris.StatusNotFound))
		ctx.View("errors/404.html")
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.View("errors/500.html")
		ctx.ViewLayout(iris.NoLayout)
	})

	// register the routes & the public API
	registerRoutes( app )
	registerAPI( app )

	// start the server
	//app.Listen("127.0.0.1:8080")
	//app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.YAML("./configs/cfg_run_app.yml")) )
	app.Run(iris.Addr(appcfg.G_GOHost))
}

/*
func initConfigration(app *iris.Application) map[string]interface{} {
	app.Configure(iris.WithConfiguration(iris.YAML("./backend/configs/cfg_run_app.yml")))
	cfg_all := app.ConfigurationReadOnly().GetOther()
	if _, ok := cfg_all["H5Host"]; ok {
		const_vars.G_H5Host = cfg_all["H5Host"].(string)
	}
	if _, ok := cfg_all["GOHost"]; ok {
		const_vars.G_GOHost = cfg_all["GOHost"].(string)
	}

	if _, ok := cfg_all["LogPath"]; ok {
		const_vars.G_LogPath = cfg_all["LogPath"].(string)
	}

	return cfg_all
}
*/

func registerRoutes( app *iris.Application ) {
	// register index using a 'Handler'
	app.Get("/", routes.Index)

	//app.Get("/-/:rand(.*).hot-update.:ext(.*)", iris.ToHandler(routes.ReloadProxy))
	//app.Get("/-/bundle.js", iris.Handler(routes.ReloadProxy))

	mvc.Configure(app.Party("/-/bundle.js"))//.Handle( routes.ReloadProxy)

	// this is other way to declare a route
	// using a 'HandlerFunc'
	app.Get("/about", routes.About)

	// Dynamic route

	//app.Get("/profile/:username", routes.Profile).ChangeName("user-profile")


	route_profile := app.Get("/profile/:username", routes.Profile)
	route_profile.Name = "user-profile"

	app.Get("/all", routes.UserList)
}

func registerAPI(  app *iris.Application  ) {
	p := app.Party("/api/users")
	{
		p.Get("/", userAPI.GetAll)
		p.Get("/:id", userAPI.GetByID)
		p.Put("/", userAPI.Insert)
		p.Post("/:id", userAPI.Update)
		p.Delete("/:id", userAPI.DeleteByID)
	}
}



// get a filename based on the date, file logs works that way the most times
// but these are just a sugar.
func todayFilename() string {
	today := time.Now().Format("2006-01-02")
	return today + ".txt"
}

func newLogFile( ) *os.File {
	filename := todayFilename()

	filefullpath := appcfg.G_LogPath + "/" + filename

	var runningMode int = 0
	/*
	cfg_main map[string]interface{}
	if _, ok := cfg_main["RuntimeMode"].(string); ok {
		if runningMode, err := strconv.Atoi(cfg_main["RuntimeMode"].(string)); err == nil {
			fmt.Printf("i=%d, type: %T\n", runningMode, runningMode)
		}
	}
	*/

	runningMode = appcfg.RuntimeMode
	// open an output file, this will append to the today's file if server restarted.
	f, err := os.OpenFile(filefullpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		if ( runningMode == const_vars.DevMode ){
			println(err)
		}else{
			panic(err)
		}

	}

	return f
}