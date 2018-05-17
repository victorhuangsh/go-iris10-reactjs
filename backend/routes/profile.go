package routes

import "github.com/kataras/iris"

func Profile(ctx iris.Context) {
	username := ctx.Params().Get("username")
	ctx.ViewData( "Username", username )
	ctx.View("profile.html")
}
