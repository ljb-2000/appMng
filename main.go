package main

import (
	_ "appMng/routers"
	"github.com/astaxie/beego"
	"appMng/utils/db"
	_ "github.com/go-sql-driver/mysql"
	_ "appMng/models"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	db.InitDatabase()
	beego.Run()
}
