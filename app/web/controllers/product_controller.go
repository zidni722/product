package controllers

import (
	"github.com/kataras/iris"
	"github.com/spf13/viper"
)

func GetHomeHandler(ctx iris.Context) {
	ctx.ViewData("Title", "Home")
	ctx.ViewData("AppName", viper.GetString("app.name"))
	ctx.ViewData("AppOwner", viper.GetString("app.owner"))
	ctx.ViewData("Message", "This is Product API index")
	ctx.View("index.html")
}