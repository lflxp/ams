package models

//pxelinux.cfg启动配置文件
type Cfg struct {
	Id 		int64 		`xorm:"autoincr"`
	Mark 		string 		`xorm:"mark varchar(40) index(index_mark) unique(u_uuid)"`
	Level 		string 		`xorm:"level" vharchar(20)`
	Types 		string 		`xorm:"types varchar(20)"`
	Columns 	string 		`xorm:"columns varchar(40) unique(u_kv)"`
	Value 		string 		`xorm:"value varchar(200) unique(u_kv)"`
	Common 		string 		`xorm:"common varchar(40)"`
}

//pxe和ks整体模板
type Config struct {
	Id 		int64 		`xorm:"autoincr"`
	Typed 		string 		`xorm:"types varchar(20)"`
	Name 		string 		`xorm:"name vharchar(40)"`
	Message 	string 		`xorm:"message LONGTEXT"`
}

//贷款信息
type Loan struct {
	Id		int64 		`xorm:"autoincr"`
	Name		string 		`xorm:"name varchar(40)"`  //姓名
	Total   	float64		//贷款额度
	Rates  		float64		//月利率
	Months  	float64 	//还款次数
	Types		string 		`xorm:"types varchar(10)"`  //还款方式
	Bank		string 		`xorm:"bank varchar(10)"`  //还款银行
	StartTime 	string		`xorm:"starttime varchar(20)"`   //开始还款时间
}

//贷款结构
type Detail struct {
	Id 		int64
	PerMonthPay 	string  //偿还本息和
	RatePay 	string  //每月还款利息
	BenJin 		string  //每月偿还本金
	UnPay 		string  //未偿还本金
	PayTime 	string  //偿还时间
	Weeks		bool 	//是否是一个星期内
}

//贷款总结构
type Data struct {
	Total 		float64  //贷款额度
	SumRates 	string //利息总和
	Everything	[]Detail
}

//用户名
type LoginUser struct {
	Id 		int64		`xorm:"pk autoincr"`
	Email 		string		`xorm:"email varchar(40) unique(u_kv)"`
	Username 	string		`xorm:"username varchar(40) unique(u_kv)"`
	Password 	string		`xorm:"password varchar(40) unique(u_kv)"`
	Common 		string		`xorm:"common varchar(40) unique(u_kv)"`
}

//登陆历史
//println(this.Ctx.Request.Referer())
//println(this.Ctx.Request.RemoteAddr)
//println(this.Ctx.Request.RequestURI)
//println(this.Ctx.Request.Host)
//println(this.Ctx.Request.Method)
//println(this.Ctx.Request.Proto)
//println(this.Ctx.Request.UserAgent())
type LoginHistory struct {
	Id 		int64		
	Username 	string		`xorm:"username varchar(40)"`
	Referer 	string		`xorm:"referer varchar(200)"`
	RemoteAddr 	string		`xorm:"remoteAddr varchar(40)"`
	RequestURI 	string		`xorm:"requestURI varchar(40)"`
	Host 		string		`xorm:"host varchar(40)"`
	Method 		string		`xorm:"mMethod varchar(40)"`
	Proto 		string		`xorm:"proto varchar(40)"`
	UserAgent 	string		`xorm:"userAgent varchar(200)"`
	InsertTime 	string		`xorm:"inserttime varchar(20)"`
}