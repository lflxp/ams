package main

import (
	_ "github.com/lflxp/ams/routers"
	_ "github.com/lflxp/ams/utils/db"
	_ "github.com/lflxp/ams/utils/cache"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.TemplateLeft = "<<"
	beego.BConfig.WebConfig.TemplateRight = ">>"
	beego.Run()
}

