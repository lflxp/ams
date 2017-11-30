package pag

import (
	. "github.com/lflxp/ams/models"
	. "github.com/lflxp/ams/utils/db"
	"fmt"
)

func HistoryPagintor(order string, offset, limit int) map[string]interface{} {
	tmp := map[string]interface{}{}
	result := make([]LoginHistory, 0)
	rrr := new(LoginHistory)
	total, _ := Db.Engine.Count(rrr)
	err := Db.Engine.Desc("Id").Limit(limit, offset).Find(&result)
	if err != nil {
		return nil
	}
	tmp["total"] = fmt.Sprintf("%d", total)
	tmp_data := []map[string]string{}
	for _, data := range result {
		tmp1 := map[string]string{}
		tmp1["id"] = fmt.Sprintf("%d", data.Id)
		tmp1["username"] = data.Username
		tmp1["referer"] = data.Referer
		tmp1["remoteaddr"] = data.RemoteAddr
		tmp1["requesturi"] = data.RequestURI
		tmp1["host"] = data.Host
		tmp1["method"] = data.Method
		tmp1["proto"] = data.Proto
		tmp1["useragent"] = data.UserAgent
		tmp1["inserttime"] = data.InsertTime
		tmp_data = append(tmp_data, tmp1)
	}
	tmp["rows"] = tmp_data
	return tmp
}

func Search(order,search string, offset, limit int) map[string]interface{} {
	tmp := map[string]interface{}{}
	result := make([]LoginHistory, 0)
	rrr := new(LoginHistory)
	total, _ := Db.Engine.Where("username = ? or referer = ? or remoteaddr = ? or host = ? or mMethod = ? or proto = ? or useragent = ?",search,search,search,search,search,search,search).Count(rrr)
	err := Db.Engine.Where("username = ? or referer = ? or remoteaddr = ? or host = ? or mMethod = ? or proto = ? or useragent = ?",search,search,search,search,search,search,search).Desc("Id").Limit(limit, offset).Find(&result)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	tmp["total"] = fmt.Sprintf("%d", total)
	tmp_data := []map[string]string{}
	for _, data := range result {
		tmp1 := map[string]string{}
		tmp1["id"] = fmt.Sprintf("%d", data.Id)
		tmp1["username"] = data.Username
		tmp1["referer"] = data.Referer
		tmp1["remoteaddr"] = data.RemoteAddr
		tmp1["requesturi"] = data.RequestURI
		tmp1["host"] = data.Host
		tmp1["method"] = data.Method
		tmp1["proto"] = data.Proto
		tmp1["useragent"] = data.UserAgent
		tmp1["inserttime"] = data.InsertTime
		tmp_data = append(tmp_data, tmp1)
	}
	tmp["rows"] = tmp_data
	return tmp
}