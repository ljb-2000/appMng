package controllers

import (
	"github.com/astaxie/beego"

	"encoding/json"
	"appMng/models"
	"log"
	"appMng/utils/git"
	"appMng/utils/ua"
	"github.com/satori/go.uuid"
	"github.com/Jeffail/gabs"
	"github.com/astaxie/beego/logs"
	"net/http"
	"time"
	"appMng/utils/k8s"
)

type AppController struct {
	beego.Controller
}

// @Title Get Apps
// @Description get apps
// @router / [get]
func (this *AppController) GetApps() {
	userId := this.Ctx.Input.Header("UserName")
	apps, err := models.GetApps(userId)
	if err != nil {
		logs.Error("获取所有应用失败 %v", err)
		this.CustomAbort(http.StatusInternalServerError, "获取所有应用失败")
	}

	k8s.GetAppsStatus(userId, apps)

	this.Data["json"] = apps
	this.ServeJSON()
}

// @Title CreateApp
// @Description create a new app
// @router / [post]
func (this *AppController) CreateApp() {
	var ob models.App
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	log.Println(ob)

	userId := this.Ctx.Input.Header("UserName")
	userName, pwd := ua.GetUserNamePwd(userId)
	log.Println("username: " + userName)
	log.Println("password: " + pwd)

	if models.IsNameUsed(userName, ob.Name) {
		logs.Error("应用名称重复")
		this.CustomAbort(http.StatusInternalServerError, "应用名称已占用")
	}

	//create a repository for the app
	giturl := git.CreateRepo(userName, pwd, ob.Name, ob.Description)
	ob.Git = giturl

	ob.Id = uuid.NewV4().String()
	ob.User = userId
	ob.State = "created"

	t := time.Now()
	ob.CreatedTime = t.Format("2006-01-02 15:04:05")

	models.AddApp(&ob)

	jsonObj := gabs.New()
	jsonObj.Set("0", "code")
	jsonObj.Set("OK", "msg")
	jsonObj.Set(ob, "data")

	this.Ctx.Output.Header("Content-Type", "application/json")
	this.Ctx.Output.Body([]byte(jsonObj.String()))
}


// @Title DeleteApp
// @Description delete the app
// @Param appId path string true "待删除的appId"
// @router /:appId [delete]
func (this *AppController) DeleteApp() {
	appId := this.Ctx.Input.Param(":appId")
	//get app name from database
	app, aerr := models.GetaApp(appId)
	if aerr != nil {
		logs.Error("获取应用失败 %v", aerr)
		this.CustomAbort(http.StatusInternalServerError, "获取应用失败")
	}

	//delete from database
	log.Println("delete app from db : " + appId)
	err := models.DeleteApp(appId)
	if err != nil {
		logs.Error("删除应用失败 %v", err)
		this.CustomAbort(http.StatusInternalServerError, "删除应用失败")
	}

	//delete from git
	userId := this.Ctx.Input.Header("UserName")
	userName, pwd := ua.GetUserNamePwd(userId)
	git.DeleteRepo(userName, pwd, app.Name)
	jsonObj := gabs.New()
	jsonObj.Set("0", "code")
	jsonObj.Set("OK", "msg")

	this.Ctx.Output.Header("Content-Type", "application/json")
	this.Ctx.Output.Body([]byte(jsonObj.String()))
}

// @Title Get An App
// @Description get an app
// @router /:appId [get]
func (this *AppController) GetAnApp() {
	appId := this.Ctx.Input.Param(":appId")
	log.Println(appId)
	app, err := models.GetaApp(appId)
	if err != nil {
		logs.Error("获取应用详情失败 %v", err)
		this.CustomAbort(http.StatusInternalServerError, "获取应用详情失败")
	}

	jsonObj := gabs.New()
	jsonObj.Set("0", "code")
	jsonObj.Set("OK", "msg")
	jsonObj.Set(app, "data")

	this.Ctx.Output.Header("Content-Type", "application/json")
	this.Ctx.Output.Body([]byte(jsonObj.String()))
}
