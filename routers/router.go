package routers

import (
	"github.com/lflxp/ams/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/login/?:type", &controllers.EasyInstallController{},"get,post:Login")
    beego.Router("/test", &controllers.MainController{},"get,post:Test")
    beego.Router("/vue", &controllers.MainController{},"get,post:Vue")
    beego.Router("/tag", &controllers.MainController{},"get,post:Tag")
    beego.Router("/", &controllers.MainController{},"get,post:Main")
    beego.Router("/config", &controllers.MainController{},"get,post:Config")
    beego.Router("/list", &controllers.MainController{},"get,post:List")
    beego.Router("/options/?:type", &controllers.MainController{},"get,post:Options")
}
