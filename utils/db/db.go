package db

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	//"github.com/astaxie/beego"
	. "github.com/lflxp/ams/models"
	. "github.com/lflxp/ams/utils/tool"
	"github.com/go-xorm/core"
	//"os"
	//"strings"
	//"fmt"
)

type Xorm struct {
	Engine *xorm.Engine
}

var (
	Db  Xorm
	err error
)

func init() {
	//Conn
	//fmt.Println(beego.AppConfig.String("db::type"),beego.AppConfig.String("db::path"))
	//Db.Engine,err = xorm.NewEngine(beego.AppConfig.String("db::type"),beego.AppConfig.String("db::path"))
	Db.Engine, err = xorm.NewEngine("sqlite3", "./db.sqlite3")
	Check(err)
	//日志 会在控制台打印出生成的SQL语句；
	Db.Engine.ShowSQL(true)
	//会在控制台打印调试及以上的信息；
	Db.Engine.Logger().SetLevel(core.LOG_OFF)
	//如果希望将信息不仅打印到控制台，而是保存为文件
	//f, err := os.Create("sql.log")
	//Check(err)
	//Db.Engine.SetLogger(xorm.NewSimpleLogger(f))
	//连接池
	Db.Engine.SetMaxIdleConns(300)
	Db.Engine.SetMaxOpenConns(300)
	//名称映射规则
	Db.Engine.SetMapper(core.SnakeMapper{})
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "pxe_")
	Db.Engine.SetTableMapper(tbMapper)
	Db.Engine.SetColumnMapper(core.SameMapper{})

	err := Db.Engine.Sync2(new(ServiceConfig),new(Service),new(CmdbTree),new(LoginHistory))
	Check(err)

	//err = Db.Engine.Sync2(new(Userinfo),new(Groups),new(Permissions))
	//Check(err)
}