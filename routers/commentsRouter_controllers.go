package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["talkGo/controllers:RoomController"] = append(beego.GlobalControllerRouter["talkGo/controllers:RoomController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/Create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["talkGo/controllers:RoomController"] = append(beego.GlobalControllerRouter["talkGo/controllers:RoomController"],
		beego.ControllerComments{
			Method: "GetList",
			Router: `/get-list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["talkGo/controllers:RoomController"] = append(beego.GlobalControllerRouter["talkGo/controllers:RoomController"],
		beego.ControllerComments{
			Method: "IExit",
			Router: `/i-exit`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["talkGo/controllers:TalkController"] = append(beego.GlobalControllerRouter["talkGo/controllers:TalkController"],
		beego.ControllerComments{
			Method: "Say",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["talkGo/controllers:TalkController"] = append(beego.GlobalControllerRouter["talkGo/controllers:TalkController"],
		beego.ControllerComments{
			Method: "MsgList",
			Router: `/msg_list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["talkGo/controllers:TalkController"] = append(beego.GlobalControllerRouter["talkGo/controllers:TalkController"],
		beego.ControllerComments{
			Method: "UpVoice",
			Router: `/upVoice`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["talkGo/controllers:UserController"] = append(beego.GlobalControllerRouter["talkGo/controllers:UserController"],
		beego.ControllerComments{
			Method: "CheckLogin",
			Router: `/check_login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["talkGo/controllers:UserController"] = append(beego.GlobalControllerRouter["talkGo/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
