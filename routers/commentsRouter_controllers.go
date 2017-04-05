package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["appMng/controllers:AppController"] = append(beego.GlobalControllerRouter["appMng/controllers:AppController"],
		beego.ControllerComments{
			Method: "GetApps",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appMng/controllers:AppController"] = append(beego.GlobalControllerRouter["appMng/controllers:AppController"],
		beego.ControllerComments{
			Method: "CreateApp",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appMng/controllers:AppController"] = append(beego.GlobalControllerRouter["appMng/controllers:AppController"],
		beego.ControllerComments{
			Method: "DeleteApp",
			Router: `/:appId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appMng/controllers:AppController"] = append(beego.GlobalControllerRouter["appMng/controllers:AppController"],
		beego.ControllerComments{
			Method: "GetAnApp",
			Router: `/:appId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appMng/controllers:AppController"] = append(beego.GlobalControllerRouter["appMng/controllers:AppController"],
		beego.ControllerComments{
			Method: "ModifyAnApp",
			Router: `/:appId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appMng/controllers:ImageController"] = append(beego.GlobalControllerRouter["appMng/controllers:ImageController"],
		beego.ControllerComments{
			Method: "CreateImage",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appMng/controllers:ImageController"] = append(beego.GlobalControllerRouter["appMng/controllers:ImageController"],
		beego.ControllerComments{
			Method: "GetImages",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appMng/controllers:ImageController"] = append(beego.GlobalControllerRouter["appMng/controllers:ImageController"],
		beego.ControllerComments{
			Method: "DeleteImage",
			Router: `/:imageId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appMng/controllers:TestController"] = append(beego.GlobalControllerRouter["appMng/controllers:TestController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
