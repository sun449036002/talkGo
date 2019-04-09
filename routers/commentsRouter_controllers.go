package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["talkGo/controllers:NotifyController"] = append(beego.GlobalControllerRouter["talkGo/controllers:NotifyController"],
        beego.ControllerComments{
            Method: "OnPlay",
            Router: `/OnPlay`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["talkGo/controllers:NotifyController"] = append(beego.GlobalControllerRouter["talkGo/controllers:NotifyController"],
        beego.ControllerComments{
            Method: "OnPlayDone",
            Router: `/OnPlayDone`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["talkGo/controllers:NotifyController"] = append(beego.GlobalControllerRouter["talkGo/controllers:NotifyController"],
        beego.ControllerComments{
            Method: "OnPublish",
            Router: `/onPublish`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["talkGo/controllers:NotifyController"] = append(beego.GlobalControllerRouter["talkGo/controllers:NotifyController"],
        beego.ControllerComments{
            Method: "OnPublishDone",
            Router: `/onPublishDone`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["talkGo/controllers:RiddleController"] = append(beego.GlobalControllerRouter["talkGo/controllers:RiddleController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/get`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["talkGo/controllers:RoomController"] = append(beego.GlobalControllerRouter["talkGo/controllers:RoomController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/Create`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["talkGo/controllers:RoomController"] = append(beego.GlobalControllerRouter["talkGo/controllers:RoomController"],
        beego.ControllerComments{
            Method: "UploadImg",
            Router: `/UploadImg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["talkGo/controllers:RoomController"] = append(beego.GlobalControllerRouter["talkGo/controllers:RoomController"],
        beego.ControllerComments{
            Method: "IExit",
            Router: `/iexit`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["talkGo/controllers:RoomController"] = append(beego.GlobalControllerRouter["talkGo/controllers:RoomController"],
        beego.ControllerComments{
            Method: "GetList",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["talkGo/controllers:TalkController"] = append(beego.GlobalControllerRouter["talkGo/controllers:TalkController"],
        beego.ControllerComments{
            Method: "Say",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["talkGo/controllers:TalkController"] = append(beego.GlobalControllerRouter["talkGo/controllers:TalkController"],
        beego.ControllerComments{
            Method: "MsgList",
            Router: `/msg_list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["talkGo/controllers:TalkController"] = append(beego.GlobalControllerRouter["talkGo/controllers:TalkController"],
        beego.ControllerComments{
            Method: "UpVoice",
            Router: `/upVoice`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["talkGo/controllers:UserController"] = append(beego.GlobalControllerRouter["talkGo/controllers:UserController"],
        beego.ControllerComments{
            Method: "CheckLogin",
            Router: `/check_login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["talkGo/controllers:UserController"] = append(beego.GlobalControllerRouter["talkGo/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
