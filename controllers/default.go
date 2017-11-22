package controllers

import (
	"strings"
	"github.com/astaxie/beego"
	"github.com/lflxp/dbui/etcd"
	"github.com/lflxp/ams/utils/cmdb"
	"github.com/lflxp/ams/models"
	"github.com/lflxp/ams/utils/tool"
	. "github.com/lflxp/ams/utils/db"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplName = "main/main.html"
}

func (this *MainController) Test() {
	this.TplName = "test.html"
}

func (this *MainController) Vue() {
	this.TplName = "vue.html"
}

func (this *MainController) Tag() {
	this.TplName = "cloud.html"
}

func (this *MainController) List() {
	this.Data["Title"] = "目录"
	this.TplName = "config/list.html"
}

func (this *MainController) Main() {
	this.TplName = "main/main.html"
}

func (this *MainController) Config() {
	st := etcd.EtcdUi{Endpoints:[]string{"localhost:2379"}}
	this.Data["Brand"] = "配置管理" //top.html 主题显示
	this.Data["Tree"] = st.GetTreeByString()
	this.Data["Column"] = etcd.GetEtcdTemplate() 
	this.Data["Title"] = "配置管理"
	this.Data["Config"] = "class='active'"
	this.TplName = "config/config.html"
}

func (this *MainController) Options() {
	types := this.Ctx.Input.Param(":type")
	if this.Ctx.Request.Method == "GET" {
		if types == "history" {
			datas := make([]models.CmdbTree, 0)
			err := Db.Engine.Find(&datas)
			if err != nil {
				this.Ctx.WriteString(err.Error())
				return
			}
			this.Data["json"] = datas
			this.ServeJSON()
		} else if types == "scan" {
			sc := new(models.ServiceConfig)
			//service := new(Service)
			ip := this.GetString("ip")
			idc := this.GetString("idc")
			_, err := Db.Engine.Where("ip = ?", ip).And("idc = ?", idc).Get(sc)
			if err != nil {
				beego.Error(err.Error())
			}
			//_,err = Db.Engine.Where("sn = ?","001").Get(service)
			//if err != nil {
			//	logs.Error(err.Error())
			//}
			result := []string{}
			if sc.Ports != "" {
				if strings.Contains(sc.Ports, "|") {
					for _, v := range strings.Split(sc.Ports, "|") {
						if tool.CommTool.ScannerPort(sc.Ip + ":" + v) {
							result = append(result, "<button class='btn btn-xs btn-success'>"+v+"</button>")
						} else {
							result = append(result, "<button class='btn btn-xs btn-danger'>"+v+"</button>")
						}
					}
				} else {
					if tool.CommTool.ScannerPort(sc.Ip + ":" + sc.Ports) {
						result = append(result, "<button class='btn btn-xs btn-success'>"+sc.Ports+"</button>")
					} else {
						result = append(result, "<button class='btn btn-xs btn-danger'>"+sc.Ports+"</button>")
					}
				}
			} else {
				result = append(result, "<button class='btn btn-xs btn-danger'>无配置</button>")
			}
			var clientStatus string
			clientStatus = "<button class='btn btn-xs btn-success'>功能暂无</button>"
			//_,found := Cached.Get(service.Catalog+"/tcp@"+ip+":"+service.Port)
			//if found  {
			//	clientStatus = "<button class='btn btn-xs btn-success'>连接正常</button>"
			//} else {
			//	clientStatus = "<button class='btn btn-xs btn-danger'>未连接</button>"
			//}

			this.Ctx.WriteString(strings.Join(result, "") + "," + clientStatus)
		}
	} else if this.Ctx.Request.Method == "POST" {
		if types == "check" {
			name := this.GetString("ids")
			//xxo := cmdb.Api.ParseData(name)
			xxo := cmdb.Api.ParseDataEtcd(name,[]string{"localhost:2379"})
			this.Data["json"] = xxo
			this.ServeJSON()
			//
			//xxo := map[string]interface{}{}
			//xxo["total"] = "200"
			//xxo["rows"] = []map[string]string{map[string]string{"id":"1","name":"item10","price":"$10"},map[string]string{"id":"1","name":"item10","price":"$10"}}
			//this.Data["json"] = xxo
			//this.ServeJSON()
		}
	}
}