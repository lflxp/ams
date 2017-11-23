package tool

import (
	"github.com/astaxie/beego"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"net"
	"crypto/md5"
	"net/http"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	//. "services/util/cache"
)

type Tool int

var CommTool Tool

func (this *Tool) CheckErr(err error) {
	if err != nil {
		beego.Error(err.Error())
		//panic(err)
	}
}

//string to json string
//25add0a|lxp|lxp|lxp@lflxp.cn|Wed Jan 11 23:47:55 2017 +0800|6 days ago|think about etcd list

func (this *Tool) StringToJson(data, ip, path string, col []string) string {
	var t []string
	for _, v := range strings.Split(data, "\n") {
		ufo := strings.Split(v, "|")
		ops := `<div class="btn-group-vertical"><button type="button" class="btn btn-primary dropdown-toggle btn-xs" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"><span class="glyphicon glyphicon-th-list" aria-hidden="true"/></button><ul class="dropdown-menu multi-level" aria-labelledby="dropdownMenu"><li><a onclick="Detail(\'` + ufo[0] + `\')"><span class="glyphicon glyphicon-wrench" aria-hidden="true"></span>详细配置</a></li><li><a onclick="Reset(\'` + ufo[0] + `\')"><span class="glyphicon glyphicon-wrench" aria-hidden="true"></span>回退</a></li></ul></div>`
		tmp := []string{"op:'" + ops + "'"}
		for k, vv := range ufo {
			tmp = append(tmp, col[k]+":'"+strings.Replace(vv, "'", "", -1)+"'")
		}
		t = append(t, "{"+strings.Join(tmp, ",")+"}")
	}
	return "[" + strings.Join(t, ",") + "]"
}

func (this *Tool) StringToJsonB(data, ip, path string, col []string) string {
	var t []string
	for _, v := range strings.Split(strings.Trim(data, "\n"), "\n") {
		//logs.Debug(v)
		ufo := strings.Split(v, ":")
		ufob := strings.Split(ufo[0], " ")
		ops := `<div class="btn-group-vertical"><button type="button" class="btn btn-primary dropdown-toggle btn-xs" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"><span class="glyphicon glyphicon-th-list" aria-hidden="true"/></button><ul class="dropdown-menu multi-level" aria-labelledby="dropdownMenu"><li><a onclick="Detail(\'` + ufob[0] + `\')"><span class="glyphicon glyphicon-wrench" aria-hidden="true"></span>详细配置</a></li><li><a onclick="Reset(\'` + ufob[0] + `\')"><span class="glyphicon glyphicon-wrench" aria-hidden="true"></span>回退</a></li></ul></div>`
		tmp := []string{"op:'" + ops + "'"}
		for k, vv := range []string{ufob[0], ufob[1], ufo[1], ufo[2]} {
			tmp = append(tmp, col[k]+":'"+strings.Replace(vv, "'", "", -1)+"'")
		}
		t = append(t, "{"+strings.Join(tmp, ",")+"}")
	}
	return "[" + strings.Join(t, ",") + "]"
}

//执行命令 返回结果
func (this *Tool) ExecCommand(commands string) string {
	// basic := "for x in {1};do load=`top -bn 1|sed -n '1p'|awk '{print $12,$13,$14}'|sed 's/[[:space:]]//g'`;cpu=`top -bn 1|sed -n '3p'|awk '{print $2,$4,$6,$8,$10,$12,$14,$16}'|sed 's/[[:space:]]/,/g'`;echo $load,$cpu;done"
	// cmd := exec.Command("bash", "-c", commands)
	// stdout, err := cmd.StdoutPipe()
	// util.CheckErr(err)
	out, err := exec.Command("bash", "-c", commands).Output()
	this.CheckErr(err)
	// bytesErr, err := ioutil.ReadAll(out)
	// util.CheckErr(err)
	return string(out)
}

//random
func (this *Tool) RandString(data string) string {
	if strings.Contains(data, ",") {
		tmp := strings.Split(data, ",")
		return tmp[int(time.Now().UnixNano())%len(tmp)]
	} else {
		return data
	}
}

//got ip
func (this *Tool) GetIps() []string {
	data := []string{}
	info, _ := net.InterfaceAddrs()
	for _, addr := range info {
		data = append(data, strings.Split(addr.String(), "/")[0])
	}
	return data
}

func (this *Tool) OpenBrower(port string) {
	this.ExecCommand("/usr/bin/x-www-browser http://127.0.0.1:" + port)
}

//判断文件是否存在
func (this *Tool) IsExistFile(file string) bool {
	_, err := os.Open(file)
	if err != nil && os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

//通过url下载文件 url->下载地址 dest->保存路径
func (this *Tool) Download(url, dest string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dest, body, 0777)
	if err != nil {
		return err
	}
	return nil
}

func (this *Tool) ScannerPort(ipAndPort string) bool {
	rs := false
	//tcpaddr,_ := net.ResolveTCPAddr("tcp4",ipAndPort)
	//_,err := net.DialTCP("tcp",nil,tcpaddr)
	_, err := net.DialTimeout("tcp", ipAndPort, 500*time.Millisecond)
	if err == nil {
		rs = true
	}
	return rs
}

func (this *Tool) StaticCmdb() string {
	return `[{"url":"http://172.16.1.71:8080/apiorg/1/?format=json","parentOrg":null,"orgleader":null,"users":["4 4"],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"思建","createdTime":"2016-11-24T08:56:25.462000Z","modifiedTime":"2016-11-24T08:56:25.462000Z","orglevel":1,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/2/?format=json","parentOrg":"思建","orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":["test2.kartor.cn _水土机房_172.31.0.1","test4.kartor.cn _陈家坪_1.2.1.3"],"name":"测试二","createdTime":"2016-11-24T08:56:44.690000Z","modifiedTime":"2016-12-10T07:27:10.643361Z","orglevel":2,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/3/?format=json","parentOrg":null,"orgleader":null,"users":["4 4"],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"思建2","createdTime":"2016-11-24T10:09:12.012000Z","modifiedTime":"2016-11-24T10:09:12.012000Z","orglevel":1,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/4/?format=json","parentOrg":null,"orgleader":null,"users":[],"wginfopy":["None_机房信息_0:水土_123"],"appinfo":[],"wginfovm":["123123123123 _水土机房_3.3.3.34","test2.kartor.cn _水土机房_172.31.0.1"],"name":"思建3","createdTime":"2016-11-24T10:09:36.380000Z","modifiedTime":"2016-12-10T07:35:43.012220Z","orglevel":1,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/5/?format=json","parentOrg":"测试二","orgleader":null,"users":["4 4","5 5","李男 12@34.cn"],"wginfopy":["www.kartor.cn_机房信息_3:水土机房_9.9.9.8"],"appinfo":["test1.kartor.cn testapp3","test1.kartor.cn testapp4","test1.kartor.cn testapp123","None testapp314","None testapp31","None testapp31123123123"],"wginfovm":["13888888888 _13_172.31.0.1"],"name":"运维名称改变","createdTime":"2016-11-24T10:53:09.863000Z","modifiedTime":"2016-12-15T06:12:31.464939Z","orglevel":3,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/6/?format=json","parentOrg":"运维名称改变","orgleader":null,"users":["gemen ge@cstonline.com","测试 12353@gmail.com","胡林 anbao789@126.com"],"wginfopy":["www.kartor.cn_机房信息_3:水土机房_2.1.1.2"],"appinfo":["test1.kartor.cn testapp1"],"wginfovm":["test1.kartor.cn _龙洲湾_1.9.9.9","test2.kartor.cn _水土机房_172.31.0.2","123123123123 _水土机房_3.3.3.34"],"name":"运维11","createdTime":"2016-11-24T10:54:17.658000Z","modifiedTime":"2016-12-10T06:31:53.015233Z","orglevel":4,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/7/?format=json","parentOrg":"运维11","orgleader":null,"users":["李男 12@34.cn"],"wginfopy":[],"appinfo":["None testapp3"],"wginfovm":["test1.kartor.cn _龙洲湾_1.9.9.9"],"name":"运维111","createdTime":"2016-11-24T10:59:20.489000Z","modifiedTime":"2016-12-02T08:08:45.163000Z","orglevel":5,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/8/?format=json","parentOrg":null,"orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"测试二二","createdTime":"2016-11-28T10:22:48.014000Z","modifiedTime":"2016-11-28T10:22:48.014000Z","orglevel":1,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/9/?format=json","parentOrg":null,"orgleader":null,"users":[],"wginfopy":["None_机房信息_0:水土_123"],"appinfo":[],"wginfovm":["123123123123 _水土机房_3.3.3.34","test2.kartor.cn _水土机房_172.31.0.1"],"name":"增加组33","createdTime":"2016-11-28T10:29:54.790000Z","modifiedTime":"2016-12-20T06:30:31.182949Z","orglevel":1,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/10/?format=json","parentOrg":null,"orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"增加组34","createdTime":"2016-11-28T10:30:41.449000Z","modifiedTime":"2016-11-30T01:38:51.073000Z","orglevel":1,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/11/?format=json","parentOrg":null,"orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"增加35","createdTime":"2016-11-28T10:31:13.421000Z","modifiedTime":"2016-11-28T10:31:13.421000Z","orglevel":1,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/12/?format=json","parentOrg":"增加35","orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"增加36","createdTime":"2016-11-28T10:33:23.257000Z","modifiedTime":"2016-11-30T02:07:48.061000Z","orglevel":2,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/13/?format=json","parentOrg":"运维11","orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"测试","createdTime":"2016-11-28T10:52:22.415000Z","modifiedTime":"2016-11-28T10:52:22.415000Z","orglevel":5,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/14/?format=json","parentOrg":"增加35","orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"增加37","createdTime":"2016-11-29T02:17:52.823000Z","modifiedTime":"2016-11-30T02:16:46.280000Z","orglevel":2,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/15/?format=json","parentOrg":"增加35","orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"增加38","createdTime":"2016-11-29T02:24:06.154000Z","modifiedTime":"2016-11-29T02:24:06.154000Z","orglevel":2,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/16/?format=json","parentOrg":"增加35","orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"增加39","createdTime":"2016-11-29T02:25:49.099000Z","modifiedTime":"2016-11-29T02:25:49.099000Z","orglevel":2,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/17/?format=json","parentOrg":null,"orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"增加40","createdTime":"2016-11-29T02:29:10.180000Z","modifiedTime":"2016-11-29T02:29:10.181000Z","orglevel":1,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/18/?format=json","parentOrg":null,"orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"增加41","createdTime":"2016-11-29T02:29:27.035000Z","modifiedTime":"2016-11-29T02:29:27.036000Z","orglevel":1,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/19/?format=json","parentOrg":"增加35","orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"增加41","createdTime":"2016-11-29T02:31:23.792000Z","modifiedTime":"2016-11-30T02:13:49.548000Z","orglevel":2,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/20/?format=json","parentOrg":"增加35","orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"增加42","createdTime":"2016-11-29T02:31:41.974000Z","modifiedTime":"2016-11-29T02:31:42.531000Z","orglevel":2,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/21/?format=json","parentOrg":"增加42","orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"增加43","createdTime":"2016-11-29T02:40:27.885000Z","modifiedTime":"2016-11-29T02:40:27.885000Z","orglevel":3,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/22/?format=json","parentOrg":"增加42","orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"增加44","createdTime":"2016-11-29T02:40:40.291000Z","modifiedTime":"2016-11-29T02:40:40.291000Z","orglevel":3,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/23/?format=json","parentOrg":null,"orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"增加52","createdTime":"2016-11-29T02:50:22.461000Z","modifiedTime":"2016-11-29T02:50:22.780000Z","orglevel":1,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/24/?format=json","parentOrg":null,"orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"增加53","createdTime":"2016-11-29T02:51:27.226000Z","modifiedTime":"2016-11-29T02:51:27.563000Z","orglevel":1,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/25/?format=json","parentOrg":null,"orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"增加54","createdTime":"2016-11-29T02:52:10.263000Z","modifiedTime":"2016-11-29T02:52:10.728000Z","orglevel":1,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/26/?format=json","parentOrg":"运维11","orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"测试2","createdTime":"2016-11-29T02:55:19.565000Z","modifiedTime":"2016-11-29T02:55:19.565000Z","orglevel":5,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/27/?format=json","parentOrg":"思建37","orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"测试21","createdTime":"2016-11-30T08:00:20.769000Z","modifiedTime":"2016-11-30T08:00:20.769000Z","orglevel":6,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/28/?format=json","parentOrg":"思建2","orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"思建22","createdTime":"2016-11-30T08:03:19.717000Z","modifiedTime":"2016-12-07T10:23:04.750000Z","orglevel":2,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/29/?format=json","parentOrg":"思建3","orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"思建31","createdTime":"2016-11-30T08:05:12.826000Z","modifiedTime":"2016-11-30T08:05:12.826000Z","orglevel":2,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/30/?format=json","parentOrg":"思建3","orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"思建32","createdTime":"2016-11-30T08:09:43.916000Z","modifiedTime":"2016-11-30T08:09:43.916000Z","orglevel":2,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/31/?format=json","parentOrg":"思建3","orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"思建43","createdTime":"2016-11-30T08:10:54.403000Z","modifiedTime":"2016-11-30T08:10:54.404000Z","orglevel":2,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/32/?format=json","parentOrg":"思建3","orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"思建35","createdTime":"2016-11-30T08:13:14.629000Z","modifiedTime":"2016-11-30T08:13:14.629000Z","orglevel":2,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/33/?format=json","parentOrg":"思建3","orgleader":null,"users":["gemen ge@cstonline.com","测试 12353@gmail.com"],"wginfopy":["asd_机房信息_2:龙洲湾_1.1.2.2"],"appinfo":[],"wginfovm":["123123123123 _水土机房_3.3.3.34","test3.kartor.cn _陈家坪_1.2.1.2"],"name":"思建36","createdTime":"2016-11-30T08:13:48.568000Z","modifiedTime":"2016-12-10T09:24:45.077087Z","orglevel":2,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/34/?format=json","parentOrg":"思建3","orgleader":null,"users":["helloworld2 hellowo@cstonline1.com","李男 12@34.cn"],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"思建37","createdTime":"2016-11-30T08:15:10.886000Z","modifiedTime":"2016-12-10T06:33:02.268441Z","orglevel":2,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/35/?format=json","parentOrg":"增加42","orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"增加77","createdTime":"2016-11-30T08:16:07.242000Z","modifiedTime":"2016-11-30T08:16:07.242000Z","orglevel":3,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/36/?format=json","parentOrg":"增加42","orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"t2","createdTime":"2016-11-30T08:18:47.677000Z","modifiedTime":"2016-11-30T08:18:47.677000Z","orglevel":3,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/37/?format=json","parentOrg":"增加42","orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":["test1.kartor.cn _龙洲湾_1.9.9.9"],"name":"t3","createdTime":"2016-11-30T08:22:19.254000Z","modifiedTime":"2016-11-30T08:22:19.254000Z","orglevel":3,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/38/?format=json","parentOrg":"增加42","orgleader":null,"users":[],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"t41","createdTime":"2016-11-30T08:28:54.030000Z","modifiedTime":"2016-12-07T02:07:39.206000Z","orglevel":3,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/39/?format=json","parentOrg":"t2","orgleader":null,"users":["helloworld2 hellowo@cstonline1.com"],"wginfopy":[],"appinfo":[],"wginfovm":[],"name":"test222","createdTime":"2016-12-01T08:21:08.816000Z","modifiedTime":"2016-12-10T06:32:54.797398Z","orglevel":4,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/40/?format=json","parentOrg":null,"orgleader":null,"users":["helloworld2 hellowo@cstonline1.com","gemen ge@cstonline.com","何小妹 dddd@cstonline.com","李好 123@gmail.com","测试 12353@gmail.com","不知道 12@126.com","胡林 anbao789@126.com"],"wginfopy":["xxd-System-Product-Name_机房信息_1:陈家坪_1.1.1.2","asd_机房信息_2:龙洲湾_1.1.2.2","www.kartor.cn_机房信息_3:水土机房_2.1.1.2"],"appinfo":[],"wginfovm":["test2.kartor.cn _水土机房_172.31.0.1","test3.kartor.cn _陈家坪_1.2.1.2","test4.kartor.cn _陈家坪_1.2.1.3","asd _水土_1.2.1.4","lxp-System-Product-Name _水土机房_172.31.0.1","ewqe _陈家坪_1.2.1.2"],"name":"t45","createdTime":"2016-12-07T10:23:55.393000Z","modifiedTime":"2016-12-10T06:32:47.280114Z","orglevel":1,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/41/?format=json","parentOrg":"思建3","orgleader":null,"users":["helloworld2 hellowo@cstonline1.com","world 1@cstonline.com","gemen ge@cstonline.com","何小妹 dddd@cstonline.com","李好 123@gmail.com","测试 12353@gmail.com","不知道 12@126.com","胡林 anbao789@126.com","helloworld3 1224343@gmail.com","何何 124dfefe@gmail.com","helloworld2 123@gmail.com","helloworld3 234@gmail.com","helloworld4 456@gmail.com","1 1","20161210 1@1.cn","33333 3","4 4","5 5"],"wginfopy":["None_机房信息_0:水土_123","xxd-System-Product-Name_机房信息_1:陈家坪_1.1.1.2","asd_机房信息_2:龙洲湾_1.1.2.2","www.kartor.cn_机房信息_3:水土机房_2.1.1.2","www.kartor.cn4_机房信息_3:水土机房_2.1.1.3","www.kartor.cn5_机房信息_3:水土机房_2.1.1.4","www.kartor.cn6_机房信息_3:水土机房_2.1.1.5","www.kartor.cn7_机房信息_3:水土机房_2.1.1.6","www.kartor.cn8_机房信息_3:水土机房_2.1.1.7","www.kartor.cn9_机房信息_3:水土机房_1.1.1.1","www.kartor.cn10_机房信息_3:水土机房_2.2.2.2","www.kartor.cn11_机房信息_3:水土机房_2.1.1.10","www.kartor.cn12_机房信息_3:水土机房_2.1.1.11","www.kartor.cn13_机房信息_3:水土机房_2.1.1.12","www.kartor.cn14_机房信息_3:水土机房_2.1.1.13","www.kartor.cn15_机房信息_3:水土机房_2.1.1.14","www.kartor.cn16_机房信息_3:水土机房_2.1.1.15","www.kartor.cn17_机房信息_3:水土机房_2.1.1.16","www.kartor.cn18_机房信息_3:水土机房_2.1.1.17","www.kartor.cn19_机房信息_3:水土机房_2.1.1.18","www.kartor.cn20_机房信息_3:水土机房_2.1.1.19","www.kartor.cn21_机房信息_3:水土机房_2.1.1.20","www.kartor.cn22_机房信息_3:水土机房_2.1.1.21","www.kartor.cn23_机房信息_3:水土机房_2.1.1.22","lxp-test.12_机房信息_3:水土机房_3.1.2.11","lxp-test.13_机房信息_3:水土机房_3.1.2.12","lxp-test.14_机房信息_3:水土机房_3.1.2.13","lxp-test.15_机房信息_3:水土机房_3.1.2.14","lxp-test.16_机房信息_3:水土机房_3.1.2.15","lxp-test.17_机房信息_3:水土机房_3.1.2.16","lxp-test.18_机房信息_3:水土机房_3.1.2.17","test.lxp-9_机房信息_1:陈家坪_3.1.1.10","test.lxp-10_机房信息_1:陈家坪_3.1.1.11","test.lxp-11_机房信息_1:陈家坪_3.1.1.12","test.lxp-12_机房信息_1:陈家坪_3.1.1.13","test.lxp-13_机房信息_1:陈家坪_3.1.1.14","test.lxp-14_机房信息_1:陈家坪_3.1.1.15","test.lxp-15_机房信息_1:陈家坪_3.1.1.16","test.lxp10_机房信息_3:水土机房_3.2.2.12","test.lxp11_机房信息_3:水土机房_3.2.2.13","test.lxp12_机房信息_3:水土机房_3.2.2.14","test.lxp13_机房信息_3:水土机房_3.2.2.15","test.lxp14_机房信息_3:水土机房_3.2.2.16","test.lxp15_机房信息_3:水土机房_3.2.2.17","test.lxp16_机房信息_3:水土机房_3.2.2.18","test.lxp17_机房信息_3:水土机房_3.2.2.19"],"appinfo":["None 我那个去","None testapp31123123123111","www.kartor.cn 魔兽世界1","www.kartor.cn 魔兽世界"],"wginfovm":["test1.kartor.cn _龙洲湾_1.9.9.9","test2.kartor.cn _水土机房_172.31.0.2","123123123123 _水土机房_3.3.3.34","test2.kartor.cn _水土机房_172.31.0.1","test3.kartor.cn _陈家坪_1.2.1.2","test4.kartor.cn _陈家坪_1.2.1.3","asd _水土_1.2.1.4","lxp-System-Product-Name _水土机房_172.31.0.1","ewqe _陈家坪_1.2.1.2","李学坪 _水土_172.31.9.99","l _水土_172.31.9.99","x _水土_172.31.9.99","p _水土_172.31.9.99","ll _水土_172.31.9.99","xx _水土_172.31.9.99","pp _水土_172.31.9.99","lll _水土_172.31.9.99","llla _水土_172.31.9.99","lllaa _水土3_172.31.9.99","lllaa17 _水土19_172.31.9.112"],"name":"SSSSSSSS","createdTime":"2016-12-10T07:36:20.647879Z","modifiedTime":"2016-12-16T03:10:45.153813Z","orglevel":2,"filed2":null,"filed3":null},{"url":"http://172.16.1.71:8080/apiorg/42/?format=json","parentOrg":"t45","orgleader":null,"users":["helloworld2 123@gmail.com","李男 12@34.cn"],"wginfopy":["lxp-life_机房信息_0:测试_127.0.0.1"],"appinfo":[],"wginfovm":["devops01.kartor.cn _洪湖西路_172.16.1.71","devops02.kartor.cn _洪湖西路_172.16.1.72","devops03.kartor.cn _洪湖西路_172.16.1.73","test1 _10_172.16.2.13","test_lxp _洪湖西路_172.20.11.25"],"name":"ansible测试","createdTime":"2016-12-21T03:18:46.712434Z","modifiedTime":"2017-01-18T11:29:17.719215Z","orglevel":2,"filed2":null,"filed3":null}]`
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func JiaMi(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}