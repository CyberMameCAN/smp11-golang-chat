package main

import (
	_ "app/routers"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func main() {
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.Run()
}
