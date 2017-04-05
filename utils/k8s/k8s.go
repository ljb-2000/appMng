package k8s

import (
	"log"
	"appMng/utils/commons"
	"encoding/json"
	"appMng/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type AppResp struct {
	Name string `json:"name"`
	Status string `json:"status"`
}

var gK8sUrl string

func init()  {
	gK8sUrl = beego.AppConfig.String("k8surl")
}

func GetAppState(user string) {
	k8sUrl := gK8sUrl + `k8s-middleware/v1/app/?namespace=` + user

	var defuser, defpwd string
	resBody, err := commons.MyTestHttpRequest("GET", k8sUrl, nil, defuser, defpwd)
	if err != nil {
		log.Println(err.Error())
	} else {

		log.Println(string(resBody))
		appResps := []AppResp{}
		json.Unmarshal(resBody, &appResps)
		log.Println(len(appResps))
	}
	return
}

func GetAppsStatus(user string, apps []models.App) error {
	k8sUrl := gK8sUrl + `k8s-middleware/v1/app/?namespace=` + user

	var defuser, defpwd string
	resBody, err := commons.MyTestHttpRequest("GET", k8sUrl, nil, defuser, defpwd)
	if err != nil {
		logs.Error("Get app status failed: %s", err.Error())
	} else {
		logs.Debug(string(resBody))
		appStatus := []AppResp{}
		json.Unmarshal(resBody, &appStatus)
		log.Println(len(appStatus))

		for i := 0; i < len(apps); i++ {
			for j := 0; j < len(appStatus); j++ {
				if apps[i].Name == appStatus[j].Name {
					apps[i].State = appStatus[j].Status
				}
			}
		}
	}
	return err
}


func GetAnAppStatus(user string, app *models.App)  {
	k8sUrl := gK8sUrl + `k8s-middleware/v1/app/?namespace=` + user

	var defuser, defpwd string
	resBody, err := commons.MyTestHttpRequest("GET", k8sUrl, nil, defuser, defpwd)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(string(resBody))
		appStatus := []AppResp{}
		json.Unmarshal(resBody, &appStatus)
		log.Println(len(appStatus))

		for j := 0; j < len(appStatus); j++ {
			if app.Name == appStatus[j].Name {
				app.State = appStatus[j].Status
			}
		}
	}
	return
}