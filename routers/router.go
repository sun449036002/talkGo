// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"talkGo/controllers"

	"github.com/astaxie/beego"
	"talkGo/controllers/cartoon"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/talk",
			beego.NSInclude(
				&controllers.TalkController{},
			),
		),
		beego.NSNamespace("/cartoon",
			beego.NSInclude(
				&cartoon.IndexController{},
			),
		),
		beego.NSNamespace("/room",
			beego.NSInclude(
				&controllers.RoomController{},
			),
		),
		beego.NSNamespace("/notify",
			beego.NSInclude(
				&controllers.NotifyController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
