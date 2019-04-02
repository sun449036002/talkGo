package controllers

import (
	"time"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/json-iterator/go"
	"github.com/astaxie/beego/httplib"
)

var intervalTimes = 2

/**
谜语
 */
type RiddleController struct {
	BaseController
}

func (c *RiddleController) URLMapping() {
	c.Mapping("Get", c.Get)
}

//谜语获取
// @Title Get
// @router /get [get]
func (c *RiddleController) Get() {
	jsonMap := make(map[string]interface{})
	//roomId := c.GetString("roomId")
	myType := c.GetString("type")

	now := time.Now()
	m := now.Minute()
	liveTimers :=  (intervalTimes - m % intervalTimes) * 60 - now.Second()

	if liveTimers > 0 && myType == "timer" {
		jsonMap["code"] = 0
		jsonMap["liveTimers"] = liveTimers

		c.Data["json"] = jsonMap
		c.ServeJSON()
		return
	}


	//API 获取谜语
	riddleUrl := "http://route.showapi.com/151-2?" + "showapi_appid=" + beego.AppConfig.String("YiYuanAppId") + "&showapi_sign=" + beego.AppConfig.String("YiYuanSecretKey")
	fmt.Println("riddleUrl is : ", riddleUrl)
	res := httplib.Get(riddleUrl)

	bts, err := res.Bytes()
	if err != nil {
		fmt.Println("err =====> ", err)
	}


	firstRiddle := jsoniter.Get(bts, "showapi_res_body", "pagebean", "contentlist").Get(0).GetInterface()

	jsonMap["code"] = 0
	jsonMap["riddle"] = firstRiddle

	c.Data["json"] = jsonMap
	c.ServeJSON()
}
