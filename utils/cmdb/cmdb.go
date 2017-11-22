package cmdb

import (
	"github.com/lflxp/dbui/etcd"
)

type Apis struct {
	Url 	string
}

var Api Apis

//获取名字为name的机构信息
//func (this *Apis) GetByName(name string) map[string]string {
//	rs := map[string]string{}
//	data := this.GetApiOrg()
//	for _,y := range *data {
//		if y["name"] == name {
//			rs = y
//			break
//		}
//	}
//	return rs
//}

//Clone
// func (this *Apis) PDTB_Clone(data ServiceConfig) string {
// 	var rs string
// 	rs = `<li class='dropdown-submenu'>
// 			<a tabindex='-1' href='javascript:;'><span class='glyphicon glyphicon-cloud' aria-hidden='true'></span>Clone 远程仓库</a>
// 			    <ul class='dropdown-menu'>`
// 	if data.GitPaths != "" {
// 		if strings.Contains(data.GitPaths,",") {
// 			for _,v := range strings.Split(data.GitPaths,",") {
// 				tmp := strings.Split(v,"/")
// 				rs += `<li><a href="#" onclick="Clone('/git/clone/?ip=`+data.Ip+"&path="+v+`')">`+tmp[len(tmp)-1]+`</a></li>`
// 			}
// 		} else {
// 			tmp := strings.Split(data.GitPaths,"/")
// 			rs += `<li><a href="#" onclick="Clone('/git/clone/?ip=`+data.Ip+"&path="+data.GitPaths+`')">`+tmp[len(tmp)-1]+`</a></li>`
// 		}
// 	} else {
// 		rs += `<li><a href="#">无配置</a></li>`
// 	}
// 	return rs + `</ul></li>`
// }

// //Pull
// func (this *Apis) PDTB_Pull(data ServiceConfig) string {
// 	var rs string
// 	rs = `<li class='dropdown-submenu'>
// 		<a tabindex='-1' href='javascript:;'><span class='glyphicon glyphicon-download' aria-hidden='true'></span>PULL到最新版本</a>
// 		    <ul class='dropdown-menu'>`
// 	if data.GitPaths != "" {
// 		if strings.Contains(data.GitPaths,",") {
// 			for _,v := range strings.Split(data.GitPaths,",") {
// 				tmp := strings.Split(v,"/")
// 				rs += `<li><a href="#" onclick="StatusGit('/git/pull/?ip=`+data.Ip+`&path=`+v+`')">`+tmp[len(tmp)-1]+`</a></li>`
// 			}
// 		} else {
// 			tmp := strings.Split(data.GitPaths,"/")
// 			rs += `<li><a href="#" onclick="StatusGit('/git/pull/?ip=`+data.Ip+`&path=`+data.GitPaths+`')">`+tmp[len(tmp)-1]+`</a></li>`
// 		}
// 	} else {
// 		rs += `<li><a href="#">无配置</a></li>`
// 	}
// 	return rs + `</ul></li>`
// }

// //Push
// func (this *Apis) PDTB_Push(data ServiceConfig) string {
// 	var rs string
// 	rs = `<li class='dropdown-submenu'>
// 		<a tabindex='-1' href='javascript:;'><span class='glyphicon glyphicon-upload' aria-hidden='true'></span>PUSH镜像</a>
// 		    <ul class='dropdown-menu'>`
// 	if data.GitPaths != "" {
// 		if strings.Contains(data.GitPaths,",") {
// 			for _,v := range strings.Split(data.GitPaths,",") {
// 				tmp := strings.Split(v,"/")
// 				rs += `<li><a href="#" onclick="StatusGit('/git/push/?ip=`+data.Ip+`&path=`+v+`')">`+tmp[len(tmp)-1]+`</a></li>`
// 			}
// 		} else {
// 			tmp := strings.Split(data.GitPaths,"/")
// 			rs += `<li><a href="#" onclick="StatusGit('/git/push/?ip=`+data.Ip+`&path=`+data.GitPaths+`')">`+tmp[len(tmp)-1]+`</a></li>`
// 		}
// 	} else {
// 		rs += `<li><a href="#">无配置</a></li>`
// 	}
// 	return rs+`</ul></li>`
// }

// //Scripts
// func (this *Apis) PDTB_Scripts(data ServiceConfig) string {
// 	var rs string
// 	rs = `<li class='dropdown-submenu'>
// 		<a tabindex='-1' href='javascript:;'><span class='glyphicon glyphicon-pushpin' aria-hidden='true'></span>自定义脚本</a>
// 		    <ul class='dropdown-menu'>`
// 	if data.Other != "" {
// 		if strings.Contains(data.Other,",") {
// 			for _,v := range strings.Split(data.Other,",") {
// 				rs += `<li><a href="#" onclick="Status('/git/scripts/?ip=`+data.Ip+`&cmd=`+v+`')">`+v+`</a></li>`
// 			}
// 		} else {
// 			rs += `<li><a href="#" onclick="Status('/git/scripts/?ip=`+data.Ip+`&cmd=`+data.Other+`')" data-toggle="modal">`+data.Other+`</a></li>`
// 		}
// 	} else {
// 		rs += `<li><a href="#">无配置</a></li>`
// 	}
// 	return rs+`</ul></li>`
// }

// //Install Client
// func (this *Apis) PDTB_Install(data ServiceConfig) string {
// 	//return `<li>
//          //       <a href="#" onclick="Install('/git/form/?ip=`+data.Ip+`')"><span class="glyphicon glyphicon-wrench" aria-hidden="true"></span>安装客户端</a>
//          //   </li>`
// 	return `<li>
//                 <a target="_blank" href="#myModal" data-toggle="modal" onclick="loadData('/git/form/?ip=`+data.Ip+`','输入`+data.Ip+`登录帐号密码')"><span class="glyphicon glyphicon-wrench" aria-hidden="true"></span>安装客户端</a>
//             </li>`
// }

// //NetStatus
// func (this *Apis) PDTB_NetStatus(data ServiceConfig) string {
// 	var rs string
// 	rs = `<li class='dropdown-submenu'>
// 		<a tabindex='-1' href='javascript:;'><span class='glyphicon glyphicon-transfer' aria-hidden='true'></span>查看网络状态</a>
// 		    <ul class='dropdown-menu'>`
// 	rs += `<li><a href="#" onclick=" Status('/config/ping?ip=`+data.Ip+`'); return false;">ping</a></li>`
// 	rs += `<li><a href="#" onclick=" Status('/config/fping?ip=`+data.Ip+`'); return false;">fping</a></li>`
// 	rs += `</ul></li>`
// 	return rs
// }

// //Web Shell
// func (this *Apis) PDTB_Shell(data ServiceConfig) string {
// 	return `<li><a href="#" onclick="WLog('/git/shell?ip=`+data.Ip+`')"><span class='glyphicon glyphicon-console' aria-hidden='true'></span>查看Console</a></li>`
// }

// //Logs  window.oper.open
// func (this *Apis) ParseDataToButton_Logs(data ServiceConfig) string {
// 	var rs string
// 	rs = `<li class='dropdown-submenu'>
// 		<a tabindex='-1' href='javascript:;'><span class='glyphicon glyphicon-console' aria-hidden='true'></span>查看运行日志</a>
// 		    <ul class='dropdown-menu'>`
// 	if data.LogPaths != "" {
// 		if strings.Contains(data.LogPaths,",") {
// 			for _,v := range strings.Split(data.LogPaths,",") {
// 				rs += `<li><a href="#" onclick="WLog('/git/logs?ip=`+data.Ip+"&logpath="+v+`')">`+v+`</a></li>`
// 			}
// 		} else {
// 			rs += `<li><a href="#" onclick="WLog('/git/logs?ip=`+data.Ip+"&logpath="+data.LogPaths+`')">`+data.LogPaths+`</a></li>`
// 		}
// 	} else {
// 		rs += `<li><a href="#">无配置</a></li>`
// 	}
// 	return rs+`</ul></li>`
// }

// //Status
// func (this *Apis) ParseDataToButton_Status(data ServiceConfig) string {
// 	var rs string
// 	rs = `<li class='dropdown-submenu'>
// 		<a tabindex='-1' href='javascript:;'><span class='glyphicon glyphicon-wrench' aria-hidden='true'></span>查看当前状态</a>
// 		    <ul class='dropdown-menu'>`
// 	if data.GitPaths != "" {
// 		if strings.Contains(data.GitPaths,",") {
// 			for _,v := range strings.Split(data.GitPaths,",") {
// 				tmp := strings.Split(v,"/")
// 				rs += `<li><a href="#" onclick="Status('/git/status?ip=`+data.Ip+"&path="+v+`')">`+tmp[len(tmp)-1]+`</a></li>`
// 			}
// 		} else {
// 			tmp := strings.Split(data.GitPaths,"/")
// 			rs += `<li><a href="#" onclick="Status('/git/status?ip=`+data.Ip+"&path="+data.GitPaths+`')">`+tmp[len(tmp)-1]+`</a></li>`
// 		}
// 	} else {
// 		rs += `<li><a href="#">无配置</a></li>`
// 	}
// 	return rs+`</ul></li>`
// }

// //Op Record
// func (this *Apis) PDTB_OpRecord(data ServiceConfig) string {
// 	var rs string
// 	rs = `<li class='dropdown-submenu'>
// 		<a tabindex='-1' href='javascript:;'><span class='glyphicon glyphicon-wrench' aria-hidden='true'></span>查看操作记录</a>
// 		    <ul class='dropdown-menu'>`
// 	if data.GitPaths != "" {
// 		if strings.Contains(data.GitPaths,",") {
// 			for _,v := range strings.Split(data.GitPaths,",") {
// 				tmp := strings.Split(v,"/")
// 				rs += `<li><a href="#" onclick="Wopen('/git/op?ip=`+data.Ip+"&path="+v+`')">`+tmp[len(tmp)-1]+`</a></li>`
// 			}
// 		} else {
// 			tmp := strings.Split(data.GitPaths,"/")
// 			rs += `<li><a href="#" onclick="Wopen('/git/op?ip=`+data.Ip+"&path="+data.GitPaths+`')">`+tmp[len(tmp)-1]+`</a></li>`
// 		}
// 	} else {
// 		rs += `<li><a href="#">无配置</a></li>`
// 	}
// 	return rs+`</ul></li>`
// }

// //Commit
// func (this *Apis) ParseDataToButton_Commit(data ServiceConfig) string {
// 	var rs string
// 	rs = `<li class='dropdown-submenu'>
// 		<a tabindex='-1' href='javascript:;'><span class='glyphicon glyphicon-wrench' aria-hidden='true'></span>查看版本记录</a>
// 		    <ul class='dropdown-menu'>`
// 	if data.GitPaths != "" {
// 		if strings.Contains(data.GitPaths,",") {
// 			for _,v := range strings.Split(data.GitPaths,",") {
// 				tmp := strings.Split(v,"/")
// 				rs += `<li><a href="#" onclick="Wopen('/git/commits?ip=`+data.Ip+"&path="+v+`')">`+tmp[len(tmp)-1]+`</a></li>`
// 			}
// 		} else {
// 			tmp := strings.Split(data.GitPaths,"/")
// 			rs += `<li><a target="_blank" onclick="Wopen('/git/commits?ip=`+data.Ip+"&path="+data.GitPaths+`')">`+tmp[len(tmp)-1]+`</a></li>`
// 		}
// 	} else {
// 		rs += `<li><a href="#">无配置</a></li>`
// 	}
// 	return rs+`</ul></li>`
// }

// //详细配置
// func (this *Apis) ParseDataToButton_Config(data ServiceConfig) string {
// 	rs := `<li><a target="_blank" href="#myModal" data-toggle="modal" onclick="loadData('`+"/jump/edit?id="+strconv.FormatInt(data.Id,10)+`','`+data.Ip+`')"><span class="glyphicon glyphicon-wrench" aria-hidden="true"></span>详细配置</a></li>`
// 	return rs
// }

// //通过数据生成操作op的button
// func (this *Apis) ParseDataToButton(data ServiceConfig) string {
// 	rs := `<div class="btn-group-vertical">
//         <button type="button" class="btn btn-primary dropdown-toggle btn-xs" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"><span class="glyphicon glyphicon-th-list" aria-hidden="true"/></button>
//         <ul class="dropdown-menu multi-level" aria-labelledby="dropdownMenu">`
// 	rs += this.ParseDataToButton_Commit(data)
// 	rs += this.PDTB_OpRecord(data)
// 	rs += this.ParseDataToButton_Status(data)
// 	rs += this.ParseDataToButton_Logs(data)
// 	rs += `<li role="separator" class="divider"></li>`
// 	rs += this.ParseDataToButton_Config(data)
// 	rs += this.PDTB_Install(data)
// 	rs += this.PDTB_Shell(data)
// 	rs += this.PDTB_NetStatus(data)
// 	rs += this.PDTB_Scripts(data)
// 	rs += `<li role="separator" class="divider"></li>`
// 	rs += this.PDTB_Pull(data)
// 	rs += this.PDTB_Push(data)
// 	rs += this.PDTB_Clone(data)
// 	return rs + `</ul></div>`
// }

// //解析api结果为map数据供json调用
// func (this *Apis) ParseData(info string) []map[string]string {
// 	config := new([]ServiceConfig)
// 	//Serviced := new(Service)
// 	//result := this.GetByName(info)
// 	err := Db.Engine.Where("groupname = ?",info).Find(config)
// 	CommTool.CheckErr(err)
// 	//_,err = Db.Engine.Where("sn = ?","001").Get(Serviced)
// 	//CommTool.CheckErr(err)
// 	//println("Service",Serviced.Catalog,Serviced.Port)
// 	xxo := []map[string]string{}
// 	for _,x := range *config {
// 		tmp := map[string]string{}
// 		tmp["主机组"] = info
// 		tmp["state"] = ""
// 		tmp["id"] = fmt.Sprintf("%d",x.Id)
// 		tmp["主机名"] = x.Hostname
// 		tmp["机房"] = x.Idc
// 		tmp["ip"] = x.Ip
// 		tmp["类型"] = x.Other
// 		tmp["API"] = x.Api
// 		tmp["ports"] = x.Ports
// 		tmp["gitpaths"] = x.GitPaths
// 		tmp["logpaths"] = x.LogPaths
// 		tmp["gitpaths"] = x.GitPaths
// 		tmp["other"] = x.Other
// 		tmp["gitpaths"] = x.GitPaths
// 		//tmp["info"] = `<button class="btn btn-xs btn-danger" data-toggle="modal" data-target="#myModal" onclick="loadData('`+"/jump/new?d="+info+","+tmp["API"]+","+x.Ip+","+x.Idc+`','`+x.Ip+`')">无配置</button>`
// 		tmp["info"] = this.ParseDataToButton(x)
// 		xxo = append(xxo,tmp)
// 	}
// 	return xxo
// }

//解析api结果为map数据供json调用
func (this *Apis) ParseDataEtcd(info string,con []string) map[string]interface{} {
	etcdd := etcd.EtcdUi{Endpoints:con}
	return etcdd.FindData(info)
}

// //只作为获取主机和组关系的api数据
// func (this *Apis) GetApiOrg() []map[string]string {
// 	var result []map[string]string
// 	data := make([]CmdbTree,0)
// 	err := Db.Engine.Find(&data)
// 	if err != nil {
// 		beego.Error(err.Error())
// 		return nil
// 	}

// 	for _,y := range data {
// 		t := map[string]string{}

// 		t["name"] = y.Name
// 		t["id"] = fmt.Sprintf("%d",y.Id)
// 		t["createdTime"] = y.Create.Format("2006-01-02 15:04:05")
// 		t["modifiedTime"] = y.Update.Format("2006-01-02 15:04:05")
// 		if y.ParentOrg == "" {
// 			t["parentOrg"] = "null"
// 		} else {
// 			t["parentOrg"] = y.ParentOrg
// 		}
// 		result = append(result,t)
// 	}

// 	return result
// }

// //获得顶级机构的信息
// func (this *Apis) GetTop(data []map[string]string) []map[string]string {
// 	rs := []map[string]string{}
// 	for _,vv := range data {
// 		if vv["parentOrg"] == "null" {
// 			rs = append(rs,vv)
// 		}
// 	}
// 	return rs
// }

// //获取所有上级机构为key的子机构（第二层，用于下面的递归）
// func (this *Apis) ForeignKeys(key string,data []map[string]string) []map[string]string {
// 	res := []map[string]string{}
// 	for _,y := range data {
// 		if y["parentOrg"] == key {
// 			res = append(res,y)
// 		}
// 	}
// 	return res
// }

// //判断是否还有子机构
// func (this *Apis) HasChild(id string,data []map[string]string) bool {
// 	ok := false
// 	for _,y := range data {
// 		if y["parentOrg"] == id {
// 			ok = true
// 		}
// 	}
// 	return ok
// }

// //递归函数，获得树形结构的关系信息 tree-view
// func (this *Apis) GetTreeRelate(top,all []map[string]string) string {
// 	result := []string{}
// 	for _,y := range top {
// 		result = append(result,"{text:'"+y["name"]+"'")
// 		if this.HasChild(y["name"],all) {
// 			result = append(result,"selectable:false,multiSelect:false,state:{expanded:false,disabled:false},'nodes':["+this.GetTreeRelate(this.ForeignKeys(y["name"],all),all)+"]}")
// 		} else {
// 			result = append(result,"icon:'glyphicon glyphicon-user',selectable:true,href:'#',ids:'"+y["name"]+"'}")
// 		}
// 	}
// 	return strings.Join(result,",")
// }

// //根据顶级机构和所有数据进行递归 得到树形结构的json字符串
// func (this *Apis) GetTreeByString() string {
// 	all := this.GetApiOrg()
// 	top := this.GetTop(all)
// 	return "["+this.GetTreeRelate(top,all)+"]"
// }