// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"appMng/controllers"

	"github.com/astaxie/beego"
)

func init() {

	nss := beego.NewNamespace("appMng/v1",
		beego.NSNamespace("/test",
			beego.NSInclude(
				&controllers.TestController{},
			),
		),
		beego.NSNamespace("/apps",
			beego.NSInclude(
				&controllers.AppController{},
			),
			beego.NSRouter("/:appId/images", &controllers.ImageController{}, "get:GetAppImages"),
		),
		beego.NSNamespace("/images",
			beego.NSInclude(
				&controllers.ImageController{},
			),
		),
	)
	beego.AddNamespace(nss)
}
