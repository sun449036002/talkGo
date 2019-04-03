package controllers

import (
	"time"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/json-iterator/go"
	"github.com/astaxie/beego/httplib"
	"talkGo/lib"
	"github.com/garyburd/redigo/redis"
	"strings"
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
	roomId := strings.Replace(c.GetString("roomId", ""), "room_", "", 1)
	myType := c.GetString("type")

	now := time.Now()
	m := now.Minute()
	liveTimers :=  (intervalTimes - m % intervalTimes) * 60 - now.Second()

	//将答案放到Redis
	rc, err := lib.Dial()
	if err != nil {
		fmt.Println(err)
	}
	defer rc.Close()

	cacheKey := "riddle_answer_" + roomId

	exists, _ := redis.Bool(rc.Do("EXISTS", cacheKey))
	if liveTimers > 0 && !exists {
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

	firstRiddle := jsoniter.Get(bts, "showapi_res_body", "pagebean", "contentlist").Get(0)
	fmt.Println("cacheKey is ", cacheKey)
	answer := firstRiddle.Get("answer").ToString() //谜底：心太软
	answerArr := strings.Split(answer, "：")
	answer = answerArr[1]
	fmt.Println(answerArr)
	_, err = rc.Do("set", cacheKey, strings.TrimSpace(answer))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("answer ===>", firstRiddle.Get("answer").ToString())
	a,_ := redis.String(rc.Do("get", cacheKey))
	//10分钟过1秒后，未答对，直接进入下一题
	rc.Do("Expire", cacheKey, 60 * intervalTimes + 1)
	fmt.Println("缓存中的答案是:", a)

	jsonMap["code"] = 0
	jsonMap["riddle"] = firstRiddle.GetInterface()

	c.Data["json"] = jsonMap
	c.ServeJSON()
}
