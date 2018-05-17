// +build bindata

package modules

import (
	"github.com/kataras/iris"
)

var (
	// serve from bindata
	Public = func(app *iris.Application) {
		// set the template engine
		//app.Adapt(view.HTML("./templates", ".html").Layout("layout.html").Binary(Asset, AssetNames))

		rv := router.NewRoutePathReverser(app, router.WithHost(const_vars.G_H5Host) )

		// locate and define our templates as usual.
		templates := iris.HTML("./templates", ".html")
		templates.Layout( "layout.html" )
		templates.Binary( Asset,AssetNames )
		app.RegisterView(templates)

		// bundle file
		app.StaticEmbedded("/bundle", "./bundle", Asset, AssetNames)

		// set static folder(s)
		app.StaticEmbedded("/public", "./public", Asset, AssetNames)
	}
)
