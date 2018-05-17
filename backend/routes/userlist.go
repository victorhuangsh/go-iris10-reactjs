package routes

import (
	"../api/user"
	"github.com/kataras/iris"
)

type (
	Page struct {
		Title string
		Users []userAPI.User
	}
)

func UserList(ctx iris.Context) {
	page := Page{"All users", userAPI.MyUsers}

	ctx.ViewData("", page )
	if err := ctx.View("userlist.html"); err != nil {
		//ctx.Log(iris.DevMode, err.Error())
		//ctx.Panic()
		ctx.Application().Logger().Errorf("error: %v", err )
		panic(err)
	}
}
