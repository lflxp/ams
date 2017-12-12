package controllers

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/astaxie/beego"
	. "github.com/lflxp/ams/models"
	"github.com/lflxp/ams/utils/cmdb"
	. "github.com/lflxp/ams/utils/db"
	"github.com/lflxp/ams/utils/tool"
	"github.com/lflxp/dbui/etcd"
	// "github.com/lflxp/ams/utils/cache"
	"github.com/lflxp/ams/utils/config"
	"github.com/lflxp/ams/utils/pag"
	"github.com/prometheus/client_golang/prometheus"
)

type MainController struct {
	beego.Controller
}

var HttpRequestCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_count",
		Help: "http request count",
	},
	[]string{"endpoint"},
)

var HttpRequestDuration = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "http_request_duration",
		Help: "http request duration",
	},
	[]string{"endpoint"},
)

//初始化默认配置管理
func init() {
	beego.Informational("初始化etcd默认信息")
	var err error
	data := map[string]string{}
	//基础路径
	data["/ams"] = "介绍:ams项目配置主路径"
	data["/ams/main"] = "介绍:主页配置路径"
	data["/ams/main/ansible"] = "ansible管理模块"
	data["/ams/main/backend"] = "后台动态标签"
	data["/ams/main/index"] = "介绍::主页网站跳转配置"
	//ansible模块
	//"获取访问端自身IP的接口"
	data["/ams/main/ansible/ip"] = fmt.Sprintf("http://%s:%s/api/v1/ip", beego.AppConfig.String("host"), beego.AppConfig.String("httpport"))
	data["/ams/main/ansible/key"] = beego.AppConfig.String("ansible::key")
	//后端动态标签
	data["/ams/main/backend/heading"] = "heading标签"
	//服务注册 拓扑图
	data["/ams/main/services"] = "介绍::服务注册及监控"
	data["/ams/main/services/server"] = "AMS系统后台::http://127.0.0.1"
	//主页配置 页面名称::跳转界面
	data["/ams/main/index/config"] = beego.AppConfig.String("tag::config")
	data["/ams/main/index/top"] = beego.AppConfig.String("tag::top")
	data["/ams/main/index/grafana"] = beego.AppConfig.String("tag::grafana")
	data["/ams/main/index/blog"] = beego.AppConfig.String("tag::blog")
	//自定义标签
	//ID::NAME::html|string
	data["/ams/main/backend/heading/2b"] = "2b::文艺青年::曾经沧海难为水 除却巫山不是云"
	st := etcd.EtcdUi{Endpoints: []string{beego.AppConfig.String("etcd::url")}}
	for key, value := range data {
		err = st.Add(key, value)
		if err != nil {
			beego.Critical(err.Error())
			return
		}
	}
	defer st.Close()
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
	start := time.Now()
	HttpRequestCount.WithLabelValues("/test").Inc()

	n := rand.Intn(100)
	if n >= 95 {
		time.Sleep(100 * time.Millisecond)
	} else {
		time.Sleep(50 * time.Millisecond)
	}

	elapsed := (float64)(time.Since(start) / time.Millisecond)
	HttpRequestDuration.WithLabelValues("/test").Observe(elapsed)
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
	var list, tab string
	Menu := map[string]string{}

	result := make([]LoginUser, 0)
	err := Db.Engine.Find(&result)
	if err != nil {
		this.Ctx.WriteString(err.Error())
	}
	//获取etcd配置信息
	st := &etcd.EtcdUi{Endpoints: []string{beego.AppConfig.String("etcd::url")}}
	st.InitClientConn()
	defer st.Close()
	resp := st.More(beego.AppConfig.String("menu::list"))

	//id::href::name::info
	for _, info := range resp.Kvs {
		if strings.ContainsAny(string(info.Value), "::") {
			tmp := strings.Split(string(info.Value), "::")
			if len(tmp) == 3 {
				l, t := config.Htmls.Create2(tmp[0], "#"+tmp[0], tmp[1], tmp[2], false)
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
	st := &etcd.EtcdUi{Endpoints: []string{beego.AppConfig.String("etcd::url")}}
	st.InitClientConn()
	defer st.Close()
	resp := st.More(beego.AppConfig.String("menu::index"))

	for _, info := range resp.Kvs {
		if strings.ContainsAny(string(info.Value), "::") {
			data = append(data, strings.Split(string(info.Value), "::")[1])
		}
	}
	this.Data["Item"] = data
	this.TplName = "main/main2.html"
}

func (this *MainController) Api() {
	types := this.Ctx.Input.Param(":type")
	if this.Ctx.Request.Method == "GET" {
		if types == "main" {
			data := map[string][]map[string]string{}
			st := &etcd.EtcdUi{Endpoints: []string{beego.AppConfig.String("etcd::url")}}
			st.InitClientConn()
			defer st.Close()
			resp := st.More(beego.AppConfig.String("menu::index"))

			for _, info := range resp.Kvs {
				if strings.ContainsAny(string(info.Value), "::") {
					tmp := map[string]string{}
					s1 := strings.Split(string(info.Value), "::")
					tmp["name"] = s1[0]
					tmp["url"] = s1[1]
					data["data"] = append(data["data"], tmp)
				}
			}
			this.Data["json"] = data
			this.ServeJSON()
		} else if types == "services" {
			data := map[string][]map[string]string{}
			st := &etcd.EtcdUi{Endpoints: []string{beego.AppConfig.String("etcd::url")}}
			st.InitClientConn()
			defer st.Close()
			resp := st.More(beego.AppConfig.String("menu::services"))

			for _, info := range resp.Kvs {
				if strings.ContainsAny(string(info.Value), "::") {
					tmp := map[string]string{}
					s1 := strings.Split(string(info.Value), "::")
					tmp["key"] = string(info.Key)
					tmp["name"] = s1[0]
					tmp["url"] = s1[1]
					data["data"] = append(data["data"], tmp)
				}
			}
			this.Data["json"] = data
			this.ServeJSON()
		} else if types == "etcd" {
			st := etcd.EtcdUi{Endpoints: []string{beego.AppConfig.String("etcd::url")}}
			rs, err := st.GetTreeByMapJtopo()
			if err != nil {
				this.Data["json"] = err.Error()
				this.ServeJSON()
				return
			}
			this.Data["json"] = rs
			this.ServeJSON()
		} else if types == "ip" {
			this.Ctx.WriteString(strings.Split(this.Ctx.Request.RemoteAddr, ":")[0])
		}
	}

}

func (this *MainController) Config() {
	types := this.Ctx.Input.Param(":type")
	if this.Ctx.Request.Method == "GET" {
		if types == "config" {
			st := etcd.EtcdUi{Endpoints: []string{beego.AppConfig.String("etcd::url")}}
			this.Data["Brand"] = "配置管理" //top.html 主题显示
			this.Data["Tree"] = st.GetTreeByString()
			this.Data["Column"] = etcd.GetEtcdServiceTemplate()
			this.Data["Title"] = "配置管理"
			this.Data["Config"] = "active"
			this.TplName = "config/config.html"
		} else if types == "top" {
			this.Data["Title"] = "全网拓扑图"
			this.Data["Top"] = "active"
			this.TplName = "config/top.html"
		}
	}
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
			var order, search string
			var offset, limit int
			this.Ctx.Input.Bind(&order, "order")
			this.Ctx.Input.Bind(&search, "search")
			this.Ctx.Input.Bind(&offset, "offset")
			this.Ctx.Input.Bind(&limit, "limit")

			if search == "" {
				this.Data["json"] = pag.HistoryPagintor(order, offset, limit)
			} else {
				this.Data["json"] = pag.Search(order, search, offset, limit)
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
			st := etcd.EtcdUi{Endpoints: []string{beego.AppConfig.String("etcd::url")}}
			err := st.Add(key, value)
			if err != nil {
				this.Ctx.WriteString(err.Error())
				return
			}
			this.Ctx.WriteString("success")
		} else if types == "delete" {
			key := this.GetString("key")
			st := etcd.EtcdUi{Endpoints: []string{beego.AppConfig.String("etcd::url")}}
			err := st.Delete(key)
			if err != nil {
				this.Ctx.WriteString(err.Error())
				return
			}
			this.Ctx.WriteString("删除成功")
		} else if types == "scan" {
			key := this.GetString("key")

			result := []string{}
			if strings.ContainsAny(key, "@") {
				tmp := strings.Split(key, "@")[1]
				if tool.CommTool.ScannerPort(tmp) {
					result = append(result, "<button class='btn btn-xs btn-success'>"+tmp+"</button>")
				} else {
					result = append(result, "<button class='btn btn-xs btn-danger'>"+tmp+"</button>")
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
			if strings.Contains(name, "ETCD->") {
				name = ""
			}
			//xxo := cmdb.Api.ParseData(name)
			xxo := cmdb.Api.ParseDataEtcd(name, []string{beego.AppConfig.String("etcd::url")})
			if strings.Contains(name, "/ams/main/services") {
				xxo["column"] = true
			} else {
				xxo["column"] = false
			}
			// beego.Critical(xxo["column"])
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
