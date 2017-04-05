package controllers

import (
	"github.com/astaxie/beego"
)

var gRegUrl string

func init() {
	gRegUrl = beego.AppConfig.String("regurl")
}
