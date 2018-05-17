// +build !bindata

package modules

import (
	"../routes"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/core/router"
	"../appconfigs"
)

var (
	Public = func( app *iris.Application ) {

		//irisMiddleware := iris.FromStd(routes.ReloadProxy)
		//app.Use(irisMiddleware)

		app.WrapRouter( routes.ReloadProxy )

		// set the template engine
		//app.Adapt(view.HTML("./frontend/templates", ".html").Layout("layout.html"))

		// user-profile is the custom,optional, route's Name: with this we can use the {{ url "user-profile" $username}} inside userlist.html

		rv := router.NewRoutePathReverser(app, router.WithHost( appconfigs.GetInst().G_GOHost ) )

		//rv := router.NewRoutePathReverser(app )
		//rv := router.NewRoutePathReverser(app, router.WithHost(const_vars.G_H5Host), router.WithScheme("http"))

		// locate and define our templates as usual.
		templates := iris.HTML("./frontend/templates", ".html")
		// add a custom func of "url" and pass the rv.URL as its template function body,
		// so {{url "routename" "paramsOrSubdomainAsFirstArgument"}} will work inside our templates.
		templates.AddFunc("url", rv.URL)

		templates.Layout( "layout.html" )
		app.RegisterView(templates)

		//app.RegisterView(iris.HTML("./frontend/templates", ".html").Layout("layout.html"))

		//app.Get("/-/:rand(.*).hot-update.:ext(.*)", iris.ToHandler(routes.ReloadProxy))
		// serve bundle file from the nodejs server looking for changes
		//app.Get("/bundle/*path", iris.ToHandler(routes.ReloadProxy))

		//mvc.Configure(app.Party("/bundle/*path")).Handle( routes.ReloadProxy)
		mvc.Configure(app.Party("/bundle/*path"))



		// set static folder(s)
		app.StaticWeb("/public", "./frontend/public")
	}
)
