package controllers

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/httplib"
	"encoding/json"
	"fmt"
	. "talkGo/models"
	. "talkGo/structs"
	"talkGo/lib"
)

type UserController struct {
	BaseController
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Login", c.Login)
	c.Mapping("CheckLogin", c.CheckLogin)
}

// Login ...
// @Title Login
// @Description get Talk by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Talk
// @Failure 403 :id is empty
// @router /login [get]
func (c *UserController) Login()  {
	code := c.GetString("code")
	userinfoJson := c.GetString("userinfo")
	fmt.Println("userinfo ===> ", userinfoJson)

	userinfo := new(Userinfo)
	err := json.Unmarshal([]byte(userinfoJson), userinfo)
	if err != nil {
		fmt.Println("userinfo json 失败")
	}

	url := beego.AppConfig.String("wxApiUrl") + "sns/jscode2session?appid=" + beego.AppConfig.String("wxSmallAppId") + "&secret=" + beego.AppConfig.String("wxSmallSecret") + "&js_code=" + code + "&grant_type=authorization_code"
	req := httplib.Get(url)
	sessionJson,err:= req.String()
	if err != nil {
		c.Data["json"] = error(err)
		c.ServeJSON()
	}

	var wxSession WxSession
	err = json.Unmarshal([]byte(sessionJson), &wxSession)
	if err != nil {
		c.Data["json"] = error(err)
		c.ServeJSON()
	}

	rc, err := lib.Dial()
	if err != nil {
		fmt.Println(err)
		c.Data["json"] = error(err)
		c.ServeJSON()
	}

	sessionCacheKey := lib.GetRandomString(16)
	sessionData, err := json.Marshal(map[string]string{"openid" : wxSession.Openid, "session_key" : wxSession.Session_key})
	if err != nil {
		fmt.Print("json.Marshal Error: ")
		fmt.Println(err)
	}
	_, err = rc.Do("SET", sessionCacheKey, string(sessionData))
	if err != nil {
		fmt.Print("redis set Error: ")
		fmt.Println(err)
	}
	fmt.Println(sessionCacheKey)

	//保存用户
	var u User
	u.Openid = wxSession.Openid
	u.Username = userinfo.NickName
	u.AvatarUrl = userinfo.AvatarUrl
	u.UserJson = userinfoJson
	err = u.NewUser()

	c.Data["json"] = map[string]string{"session_key" : sessionCacheKey}
	c.ServeJSON()
}

// CheckLogin ...
// @Title CheckLogin
// @Description get Talk by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Talk
// @Failure 403 :id is empty
// @router /check_login [get]
func (c *UserController) CheckLogin()  {
	session_key := c.GetString("session_key")
	rc, err := lib.Dial()
	if err != nil {
		fmt.Println(err)
		c.Data["json"] = error(err)
		c.ServeJSON()
	}

	sv, err := redis.String(rc.Do("get", session_key))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(session_key, "`s value is  ==>", sv)

	c.Data["json"] = map[string]string{"session_key" : string(sv)}
	c.ServeJSON()
}
