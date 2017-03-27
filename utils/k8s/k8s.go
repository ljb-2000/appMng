package k8s

import (
	"log"
	"appMng/utils/commons"
	"encoding/json"
	"appMng/models"
)

type AppResp struct {
	Name string `json:"name"`
	Status string `json:"status"`
}

func GetAppState(user string) {
	k8sUrl := `http://172.16.5.245:8080/v1/app/?namespace=` + user

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

func GetAppsStatus(user string, apps []models.App)  {
	k8sUrl := `http://172.16.5.245:8080/v1/app/?namespace=` + user

	var defuser, defpwd string
	resBody, err := commons.MyTestHttpRequest("GET", k8sUrl, nil, defuser, defpwd)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(string(resBody))
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
	return
}