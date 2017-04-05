package controllers

import (
	"github.com/astaxie/beego"

	"appMng/models"
)

type TestController struct {
	beego.Controller
}

// @Title Get
// @Description 测试router
// @router / [get]
func (this *TestController) Get() {
	this.Data["json"] = models.PrintHello()
	this.ServeJSON()
}

