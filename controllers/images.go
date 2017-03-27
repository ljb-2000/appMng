package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"log"
	"github.com/Jeffail/gabs"
	//"appMng/utils/image"
	"github.com/satori/go.uuid"
	"appMng/models"
	"time"
	"appMng/utils/tpl"
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

	if IsImageNameTagUsed(ob) {

	}

	/*
	imgAddr := image.BuildImage(ob.User, ob.Name, ob.Tag, ob.Git, ob.Lang)
	image.PushImageBack(imgAddr)
	*/

	go tpl.GenerateShellFile(ob)

	imgAddr := "registry.time-track.cn:8443/" + ob.User + "/" + ob.Name + ":" + ob.Tag
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

	/*
	imgs := []models.Image{}
	imgs, _ = models.GetImages(appId)

	jsonObj := gabs.New()
	jsonObj.Set("0", "code")
	jsonObj.Set("OK", "msg")
	jsonObj.Set(imgs, "imgs")

	this.Ctx.Output.Header("Content-Type", "application/json")
	this.Ctx.Output.Body([]byte(jsonObj.String()))
	*/

	this.Data["json"] = map[string]interface{}{"delete": imageId}
	this.ServeJSON()
}

func (this *ImageController) GetAppImages() {
	appId := this.Ctx.Input.Param(":appId")
	imgs := []models.Image{}
	imgs, _ = models.GetImages(appId)

	this.Data["json"] = imgs
	this.ServeJSON()
}