package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"log"
	"github.com/Jeffail/gabs"
	"github.com/satori/go.uuid"
	"appMng/models"
	"time"
	"appMng/utils/tpl"
	"github.com/astaxie/beego/logs"
	"net/http"
)

type ImageController struct {
	beego.Controller
}

// @Title Create Image
// @Description create a new image
// @router / [post]
func (this *ImageController) CreateImage() {

	var ob models.Image
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	ob.User = this.Ctx.Input.Header("UserName")
	ob.Id = uuid.NewV4().String()
	log.Println(ob)

	app, aerr := models.GetaApp(ob.AppId)
	if aerr != nil {
		logs.Error("获取应用失败 %v", aerr)
		this.CustomAbort(http.StatusInternalServerError, "获取应用失败")
	}

	go tpl.GenerateShellFile(ob, app.Name)

	imgAddr := "registry.time-track.cn:8052/" + ob.User + "/" + ob.Name + ":" + ob.Tag
	ob.Img = imgAddr

	t := time.Now()
	ob.CreatedTime = t.Format("2006-01-02 15:04:05")

	ob.State = "building"

	log.Println(ob)
	err := models.AddImage(&ob)
	if err != nil {
		log.Println(err.Error())
	}

	jsonObj := gabs.New()
	jsonObj.Set("0", "code")
	jsonObj.Set("OK", "msg")
	jsonObj.Set(imgAddr, "imgAddr")

	this.Ctx.Output.Header("Content-Type", "application/json")
	this.Ctx.Output.Body([]byte(jsonObj.String()))
}

// @Title Get Images
// @Description get images
// @router / [get]
func (this *ImageController) GetImages() {
	var appId string
	imgs := []models.Image{}
	imgs, _ = models.GetImages(appId)

	jsonObj := gabs.New()
	jsonObj.Set("0", "code")
	jsonObj.Set("OK", "msg")
	jsonObj.Set(imgs, "imgs")

	this.Ctx.Output.Header("Content-Type", "application/json")
	this.Ctx.Output.Body([]byte(jsonObj.String()))
}


// @Title delete Image
// @Description delete image
// @router /:imageId [delete]
func (this *ImageController) DeleteImage() {
	imageId := this.Ctx.Input.Param(":imageId")
	log.Println(imageId)

	err := models.DeleteImage(imageId)
	if err != nil {
		logs.Error("删除镜像失败 %v", err)
		this.CustomAbort(http.StatusInternalServerError, "删除镜像失败")
	}
	jsonObj := gabs.New()
	jsonObj.Set("0", "code")
	jsonObj.Set("OK", "msg")

	this.Ctx.Output.Header("Content-Type", "application/json")
	this.Ctx.Output.Body([]byte(jsonObj.String()))
}

func (this *ImageController) GetAppImages() {
	appId := this.Ctx.Input.Param(":appId")
	imgs := []models.Image{}
	imgs, _ = models.GetImages(appId)

	this.Data["json"] = imgs
	this.ServeJSON()
}