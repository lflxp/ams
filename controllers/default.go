package controllers

import (
	"fmt"
	"strings"
	// "time"
	"github.com/astaxie/beego"
	"github.com/lflxp/dbui/etcd"
	"github.com/lflxp/ams/utils/cmdb"
	. "github.com/lflxp/ams/models"
	"github.com/lflxp/ams/utils/tool"
	. "github.com/lflxp/ams/utils/db"
	// "github.com/lflxp/ams/utils/cache"
	"github.com/lflxp/ams/utils/pag"
	"github.com/lflxp/ams/utils/config"
)

type MainController struct {
	beego.Controller
}

// func (this *MainController) Prepare() {
// 	//记录访问日志
// 	this.EnableXSRF = false
// 	beegoSessionId := this.Ctx.GetCookie("beegosessionID")
// 	if _, isExist := cache.Cached.Get(beegoSessionId); isExist == false {
// 		this.Ctx.Redirect(301, "/login/login")
// 		return
// 	}
// 	//记录访问日志
// 	username, isExist := cache.Cached.Get(beegoSessionId)
// 	history := new(LoginHistory)
// 	history.InsertTime = time.Now().Format("2006-01-02 15:04:05")
// 	if isExist {
// 		history.Username = username.(string)
// 	} else {
// 		history.Username = "未登陆用户"
// 	}
// 	history.Referer = this.Ctx.Request.Referer()
// 	history.RemoteAddr = this.Ctx.Request.RemoteAddr
// 	history.RequestURI = this.Ctx.Request.RequestURI
// 	history.Host = this.Ctx.Request.Host
// 	history.Method = this.Ctx.Request.Method
// 	history.Proto = this.Ctx.Request.Proto
// 	history.UserAgent = this.Ctx.Request.UserAgent()
// 	_,err := Db.Engine.Insert(history)	
// 	if err != nil {
// 		fmt.Println("insert",err.Error())
// 	}
// }

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

/**
动态标签
ETCD_URL /ams/main/backend/heading
KEY /ams/main/backend/heading/2b
Value 2b::文艺青年::曾经沧海难为水 除却巫山不是云
说明: value由三个字段组成 分别对应的是html的:ID:NAME:TABINFO
*/
func (this *MainController) List() {
	//动态生成menu
	var list,tab string
	Menu := map[string]string{}
	
	result := make([]LoginUser, 0)
	err := Db.Engine.Find(&result)
	if err != nil {
		this.Ctx.WriteString(err.Error())
	}
	//获取etcd配置信息
	st := &etcd.EtcdUi{Endpoints:[]string{beego.AppConfig.String("etcd::url")}}
	st.InitClientConn()
	defer st.Close()
	resp := st.More(beego.AppConfig.String("menu::list"))
	
	//id::href::name::info
	for _,info := range resp.Kvs {
		if strings.ContainsAny(string(info.Value),"::") {
			tmp := strings.Split(string(info.Value),"::")
			if len(tmp) == 3 {
				l,t := config.Htmls.Create2(tmp[0],"#"+tmp[0],tmp[1],tmp[2],false)
				list += l
				tab += t
			}
		}
	}
	Menu["list"] = list
	Menu["tab"] = tab
	this.Data["List"] = "active"
	this.Data["Maps"] = Menu
	this.Data["Title"] = "后台管理"
	this.Data["Result"] = result
	this.TplName = "config/list.html"
}

/**
//主页跳转
ETCD_URL /ams/main/index
KEY /ams/main/index/config
Value 配置管理::<a href="/config"><button class="btn btn-success">配置管理</button></a>
说明： NAME:HTML的标签
*/
func (this *MainController) Main() {
	data := []string{}
	st := &etcd.EtcdUi{Endpoints:[]string{beego.AppConfig.String("etcd::url")}}
	st.InitClientConn()
	defer st.Close()
	resp := st.More(beego.AppConfig.String("menu::index"))
	
	for _,info := range resp.Kvs {
		if strings.ContainsAny(string(info.Value),"::") {
			data = append(data,strings.Split(string(info.Value),"::")[1])
		}
	}
	this.Data["Item"] = data
	this.TplName = "main/main.html"
}

func (this *MainController) Config() {
	st := etcd.EtcdUi{Endpoints:[]string{beego.AppConfig.String("etcd::url")}}
	this.Data["Brand"] = "配置管理" //top.html 主题显示
	this.Data["Tree"] = st.GetTreeByString()
	this.Data["Column"] = etcd.GetEtcdTemplate() 
	this.Data["Title"] = "配置管理"
	this.Data["Config"] = "active"
	this.TplName = "config/config.html"
}

func (this *MainController) Admin() {
	types := this.Ctx.Input.Param(":type")
	if this.Ctx.Request.Method == "GET" {
		if types == "useradd" {
			result := make([]LoginUser, 0)
			err := Db.Engine.Find(&result)
			if err != nil {
				this.Ctx.WriteString(err.Error())
			}
			this.Data["Result"] = result
			this.TplName = "admin/user/user.html"	
		} else if types == "history" {
			this.Data["Brand"] = "后台管理" //top.html 主题显示
			this.TplName = "admin/history/history.html"
		} else if types == "gethistory" {
			var order,search string
			var offset, limit int
			this.Ctx.Input.Bind(&order, "order")
			this.Ctx.Input.Bind(&search, "search")
			this.Ctx.Input.Bind(&offset, "offset")
			this.Ctx.Input.Bind(&limit, "limit")

			if search == "" {
				this.Data["json"] = pag.HistoryPagintor(order, offset, limit)
			} else {
				this.Data["json"] = pag.Search(order,search, offset, limit)
			}
			
			this.ServeJSON()
		} 
	} else {
		if types == "userdel" {
			var ids string
			this.Ctx.Input.Bind(&ids, "ids")
			data := new(LoginUser)
			if strings.Contains(ids, ",") != true {
				_, err := Db.Engine.Id(ids).Delete(data)
				if err != nil {
					this.Ctx.WriteString(err.Error())
				}
			} else {
				for _, i := range strings.Split(ids, ",") {
					_, err := Db.Engine.Id(i).Delete(data)
					if err != nil {
						this.Ctx.WriteString(err.Error())
					}
				}
			}
			this.Ctx.WriteString("删除成功")
		} else if types == "userchange" {
			var name, value, pk string
			this.Ctx.Input.Bind(&name, "name")
			this.Ctx.Input.Bind(&value, "value")
			this.Ctx.Input.Bind(&pk, "pk")

			loan := new(LoginUser)
			has, err := Db.Engine.Id(pk).Get(loan)
			if err != nil {
				this.Ctx.WriteString(err.Error())
			}
			if has == false {
				this.Ctx.WriteString("not exist")
			}
			switch name {
			case "email":
				loan.Email = value
			case "username":
				loan.Username = value
			case "password":
				loan.Password = tool.JiaMi(value)
			case "common":
				loan.Common = value
			}
			affected, err := Db.Engine.Id(pk).Update(loan)
			if err != nil {
				fmt.Println(err.Error())
				this.Ctx.WriteString(err.Error())
			}
			this.Ctx.WriteString(fmt.Sprintf("update %d success", affected))
		} else if types == "historydel" {
			var ids string
			this.Ctx.Input.Bind(&ids, "ids")
			data := new(LoginHistory)
			if strings.Contains(ids, ",") != true {
				_, err := Db.Engine.Id(ids).Delete(data)
				if err != nil {
					this.Ctx.WriteString(err.Error())
				}
			} else {
				for _, i := range strings.Split(ids, ",") {
					_, err := Db.Engine.Id(i).Delete(data)
					if err != nil {
						this.Ctx.WriteString(err.Error())
					}
				}
			}
			this.Ctx.WriteString("删除成功")
		}
	}	
}

func (this *MainController) Options() {
	types := this.Ctx.Input.Param(":type")
	if this.Ctx.Request.Method == "GET" {
		if types == "history" {
			datas := make([]CmdbTree, 0)
			err := Db.Engine.Find(&datas)
			if err != nil {
				this.Ctx.WriteString(err.Error())
				return
			}
			this.Data["json"] = datas
			this.ServeJSON()
		} else if types == "aedit" {
			key := this.GetString("key")
			value := this.GetString("value")
			st := etcd.EtcdUi{Endpoints:[]string{beego.AppConfig.String("etcd::url")}}
			err := st.Add(key,value)
			if err != nil {
				this.Ctx.WriteString(err.Error())
				return
			}
			this.Ctx.WriteString("success")
		} else if types == "delete" {
			key := this.GetString("key")
			st := etcd.EtcdUi{Endpoints:[]string{beego.AppConfig.String("etcd::url")}}
			err := st.Delete(key)
			if err != nil {
				this.Ctx.WriteString(err.Error())
				return
			}
			this.Ctx.WriteString("删除成功")
		} else if types == "scan" {
			sc := new(ServiceConfig)
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
			if strings.Contains(name,"ETCD->") {
				name = ""
			}
			//xxo := cmdb.Api.ParseData(name)
			xxo := cmdb.Api.ParseDataEtcd(name,[]string{beego.AppConfig.String("etcd::url")})
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