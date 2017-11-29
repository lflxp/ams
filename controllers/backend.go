package controllers

import (
	. "github.com/lflxp/ams/models"
	"github.com/lflxp/ams/utils/cache"
	. "github.com/lflxp/ams/utils/db"
	"github.com/lflxp/ams/utils/tool"
	"github.com/astaxie/beego"
	"html/template"
	"time"
)

type EasyInstallController struct {
	beego.Controller
}

// func (this *EasyInstallController) Prepare() {
// 	//记录访问日志
// 	beegoSessionId := this.Ctx.GetCookie("beegosessionID")
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
// 	Db.Engine.Insert(history)
// }

func (this *EasyInstallController) Login() {
	types := this.Ctx.Input.Param(":type")
	if this.Ctx.Request.Method == "GET" {
		if types == "login" {
			this.XSRFExpire = 7200
			this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
			this.TplName = "login.html"
		} else if types == "logout" {
			beegoSessionId := this.Ctx.GetCookie("beegosessionID")
			cache.Cached.Delete(beegoSessionId)
			this.Ctx.Redirect(301, "/login/login")
		}
	} else {
		if types == "register" {
			var email, username, password, password2 string
			this.Ctx.Input.Bind(&email, "email")
			this.Ctx.Input.Bind(&username, "username")
			this.Ctx.Input.Bind(&password, "password")
			this.Ctx.Input.Bind(&password2, "password2")
			beego.Critical(email, username, password, password2)
			if email != "169471087@qq.com" {
				this.Ctx.Redirect(301, "/login/login")
				return
			}
			
			user := new(LoginUser)
			has, err := Db.Engine.Where("email = ? and username = ?", email, username).Get(user)
			if err != nil {
				this.Ctx.WriteString(err.Error())
				this.Abort("401")
				return
			}
			if has {
				this.Ctx.WriteString("user Exist")
				this.Abort("401")
				return
			}
			if password == password2 {
				user.Id = time.Now().Unix()
				user.Email = email
				user.Username = username
				user.Password = tool.JiaMi(password)
				_, err := Db.Engine.Insert(user)
				if err != nil {
					this.Ctx.WriteString(err.Error())
					this.Abort("401")
					return
				}
				this.Ctx.Redirect(301, "/")
			} else {
				this.Ctx.WriteString("两次密码不一致")
			}
		} else if types == "login" {
			beegoSessionId := this.Ctx.GetCookie("beegosessionID")
			var username, password string
			this.Ctx.Input.Bind(&username, "username")
			this.Ctx.Input.Bind(&password, "password")

			user := new(LoginUser)
			has, err := Db.Engine.Where("username = ? and password = ?", username, tool.JiaMi(password)).Get(user)
			if err != nil {
				this.Ctx.Redirect(301, "/login/login")
				//this.Ctx.WriteString(err.Error())
				return
			}
			if has {
				exp, _ := time.ParseDuration("24h")
				_, isexist := cache.Cached.Get(beegoSessionId)
				if isexist == false {
					cache.Cached.Set(beegoSessionId, username, exp)
				}
			}
			//this.Ctx.Redirect(301, "/login/login")
			this.Ctx.Redirect(301, "/")
		}
	}
}
