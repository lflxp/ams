package models

import "time"

type ServiceConfig struct {
	Id 			int64		`xorm:"autoincr"`
	GroupName 		string 		`xorm:"groupname varchar(40) index(org_index)"`
	Api 			string 		`xorm:"api varchar(100)"`
	Ip 			string 		`xorm:"ip varchar(15) index(ip_index)"`
	Idc 			string 		`xorm:"idc varchar(40)"`
	Hostname 		string 		`xorm:"hostname varchar(42)"`
	Ports 			string 		`xorm:"ports varchar(200)"`
	GitPaths 		string 		`xorm:"gitpaths varchar(1000)"`
	LogPaths 		string 		`xorm:"logpaths varchar(1000)"`
	Other			string 		`xorm:"other varchar(1000)"`
}

type Service struct {
	Id 			int64 		`xorm:"autoincr"`
	EtcdUrl			string 		`xorm:"etcdurl varchar(100)"`
	Catalog 		string 		`xorm:"catalog varchar(200) unique(index_unique_c)"`
	Port 			string 		`xorm:"port varchar(20)"`
	Sn 			string 		`xorm:"sn varchar(40) index(index_sn)"`
	Common 			string 		`xorm:"common varchar(40)"`
}

//记录树形结构
type CmdbTree struct {
	Id 			int64 		`xorm:"autoincr`
	ParentOrg		string 		`xorm:"parentOrg varchar(88)`
	Name 			string 		`xorm:"name varchar(88)`
	Machine 		string		`xorm:"machine TEXT"`           //hostname1:location:ip:type,hostname2:location:ip:type
	Create     		time.Time 	`xorm:"created"`                //这个Field将在Insert时自动赋值为当前时间
	Update     		time.Time 	`xorm:"updated"`                //这个Field将在Insert或Update时自动赋值为当前时间
}
