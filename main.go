package main

import (
	_ "github.com/lflxp/ams/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.TemplateLeft = "<<"
	beego.BConfig.WebConfig.TemplateRight = ">>"
	beego.Run()
}

