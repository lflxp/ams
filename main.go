package main

import (
	_ "github.com/lflxp/ams/routers"
	_ "github.com/lflxp/ams/utils/db"
	_ "github.com/lflxp/ams/utils/cache"
	"github.com/astaxie/beego"
)

func main() {
	//设置session
	// beego.BConfig.WebConfig.Session.SessionOn = true
	// beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 7 * 24 * 3600
	// beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 7 * 24 * 3600
	beego.SetStaticPath("/img","img")
	beego.SetStaticPath("/config/img","img")
	//修改模板关键字
	beego.BConfig.WebConfig.TemplateLeft = "<<"
	beego.BConfig.WebConfig.TemplateRight = ">>"
	beego.Run()
}

