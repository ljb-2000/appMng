package ua

import (
	"github.com/astaxie/beego"
	"log"
	"appMng/utils/commons"
	"encoding/json"
	"strings"
	"github.com/jmoiron/jsonq"
)

var gUaUrl string

func init()  {
	gUaUrl = beego.AppConfig.String("uaurl")
}

func GetUserNamePwd(user string) (name, pwd string) {
	uaapi := gUaUrl + `cp-ua/v1/user/` + user
	var defuser, defpwd string
	resBody, err := commons.MyTestHttpRequest("GET", uaapi, nil, defuser, defpwd)

	if err != nil {
		log.Println(err.Error())
		log.Println("error in ua")
	} else {
		data := map[string]interface{}{}
		dec := json.NewDecoder(strings.NewReader(string(resBody)))
		dec.Decode(&data)
		jq := jsonq.NewQuery(data)
		name, _ = jq.String("user", "name")
		pwd, _ = jq.String("user", "no_enc_pwd")
	}
	return
}
