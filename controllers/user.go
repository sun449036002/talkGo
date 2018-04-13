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
	"github.com/json-iterator/go"
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
	rawData := c.GetString("rawData")
	signature := c.GetString("signature")
	encryptedData := c.GetString("encryptedData")
	iv := c.GetString("iv")
	fmt.Println("code=", code)
	fmt.Println("rawData=", rawData)
	fmt.Println("signature=", signature)
	fmt.Println("encryptedData=", encryptedData)
	fmt.Println("iv=", iv)

	url := beego.AppConfig.String("wxApiUrl") + "sns/jscode2session?appid=" + beego.AppConfig.String("wxSmallAppId") + "&secret=" + beego.AppConfig.String("wxSmallSecret") + "&js_code=" + code + "&grant_type=authorization_code"
	req := httplib.Get(url)
	sessionJson,err:= req.String()
	fmt.Println("sessionJson =====> ", sessionJson)
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

	//用户的信息
	userinfo := new(Userinfo)
	err = json.Unmarshal([]byte(rawData), userinfo)
	if err != nil {
		fmt.Println("userinfo json 失败")
	}

	//保存用户
	var u User
	u.Openid = wxSession.Openid
	u.Username = userinfo.NickName
	u.AvatarUrl = userinfo.AvatarUrl
	u.UserJson = rawData
	err = u.NewUser()

	userJson, _ := jsoniter.MarshalToString(u)
	_, err = rc.Do("SET", "userinfo_" + sessionCacheKey, userJson)
	if err != nil {
		fmt.Print("redis set Error: ")
		fmt.Println(err)
	}

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

	resultMap := make(map[string]string)
	if sv != "" {
		resultMap["session_key"] = sv

		wxSession := &WxSession{}
		err := jsoniter.UnmarshalFromString(sv, &wxSession)
		if err != nil {
			fmt.Println(err)
		}

		var u User
		u.GetUserByOpenid(wxSession.Openid)

		userJson, _ := jsoniter.MarshalToString(u)
		_, err = rc.Do("SET", "userinfo_" + session_key, userJson)
		if err != nil {
			fmt.Print("redis set Error: ")
			fmt.Println(err)
		}

		resultMap["nickname"] = u.Username
	}

	c.Data["json"] = resultMap
	c.ServeJSON()
}
