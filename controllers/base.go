package controllers

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"encoding/json"
	. "talkGo/models"
	. "talkGo/structs"
	"talkGo/lib"
)

type BaseController struct {
	beego.Controller
	user User
}

func (c *BaseController) init()  {
	//用户标识
	sessionKey := c.GetString("sk")
	rc, err := lib.Dial()
	if err != nil {
		fmt.Println(err)
	}
	sv, err := redis.String(rc.Do("GET", sessionKey))
	if err != nil {
		fmt.Println(err)
	}
	if sv != "" {
		var wxs WxSession
		err = json.Unmarshal([]byte(sv), &wxs)
		if err != nil {
			fmt.Println(err)
		}

		var u User
		u.GetUserByOpenid(wxs.Openid)
		c.user = u
	} else {
		fmt.Println(sessionKey, "sv 为这空", sv)
	}
}
