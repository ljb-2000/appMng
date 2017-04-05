package git

import (
	"github.com/astaxie/beego"
	"log"
	"encoding/json"
	"strings"
	"appMng/utils/commons"
	"github.com/jmoiron/jsonq"
	"os/exec"
	"os"
	"github.com/astaxie/beego/logs"
)


type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

var gGogsUrl string

func init()  {
	gGogsUrl = beego.AppConfig.String("gogsurl")
}

func CreateRepo(user, pwd, appname, desc string) (giturl string) {
	r := Repository{}
	r.Name = appname
	r.Description = desc
	r.Private = false
	gogsapi := gGogsUrl + `api/v1/user/repos`
	bytes, merr := json.Marshal(r)
	if merr != nil {
		logs.Error(merr.Error())
	} else {
		reqBody := strings.NewReader(string(bytes))
		resBody, err := commons.MyTestHttpRequest("POST", gogsapi, reqBody, user, pwd)
		if err != nil {
			log.Println(err.Error())
		} else {
			data := map[string]interface{}{}
			dec := json.NewDecoder(strings.NewReader(string(resBody)))
			dec.Decode(&data)
			jq := jsonq.NewQuery(data)
			giturl, _ = jq.String("clone_url")
		}
	}
	return
}

func DeleteRepo(user, pwd, appname string) {
	delrepo := gGogsUrl + `api/v1/repos/` + user + `/` + appname
	resBody, err := commons.MyTestHttpRequest("DELETE", delrepo, nil, user, pwd)
	if err != nil {
		log.Println("delete repos failed.")
		logs.Error(err.Error())
	} else {
		log.Println("delete repos succeed: " + string(resBody))
	}
}

func CloneRepo(giturl string) {
	os.Chdir("/user/")
	cmd := "git"
	args := []string{"clone", giturl}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		log.Println(err.Error())
	}
	log.Println("git clone finished.")
}