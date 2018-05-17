package routes

import (
	"github.com/kataras/iris"
)

type index struct {
	Title   string
	Message string
}

func Index(ctx iris.Context) {
	// MustRender, same as Render but sends status 500 internal server error if rendering failed
	//ctx.View("index.html", nil)
	ctx.View("index.html")
}

func (i *index) Serve(ctx iris.Context) {
	ctx.ViewData( "",i )
	if err := ctx.View("index.html"); err != nil {
		// ctx.EmitError(iris.StatusInternalServerError) =>
		//ctx.Panic()
		panic(err)
	}
}
