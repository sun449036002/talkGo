package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/httplib"
	"encoding/json"
	"encoding/base64"
	"fmt"
	"strconv"
	"math/rand"
	"os"
	"time"
	. "talkGo/models"
	. "talkGo/structs"
	"talkGo/lib"
	"log"
	"os/exec"
	"io/ioutil"
)

// TalkController operations for Talk
type TalkController struct {
	beego.Controller
}

// URLMapping ...
func (c *TalkController) URLMapping() {
	/*c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)*/
	c.Mapping("Say", c.Say)
	c.Mapping("Login", c.Login)
	c.Mapping("CheckLogin", c.CheckLogin)
	c.Mapping("UpVoice", c.UpVoice)
}

// UpVoice...
// @Title UpVoice
// @Description up voice to server,chnage to text
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Talk
// @Failure 403 :id is empty
// @router /upVoice [post]
func (c *TalkController) UpVoice() {
	f,_,err := c.GetFile("file")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()

	silkFileName := "talk_" + time.Now().Format("20060102150405") + lib.GetRandomString(3)
	ferr := c.SaveToFile("file", "static/" + silkFileName  + ".silk") // 保存位置在 static/upload, 没有文件夹要先创建
	if ferr != nil {
		fmt.Println(ferr)
	}

	//创建获取命令输出管道
	//fmt.Println("/root/go/src/talkGo/static/" + silkFileName)
	cmd := exec.Command("sh", "/root/silk-v3-decoder/converter.sh",  "/root/go/src/talkGo/static/" + silkFileName  + ".silk", "pcm")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return
	}
	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return
	}
	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
		return
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return
	}
	fmt.Printf("stdout:\n\n %s", bytes)

	//BAI DU TOKEN
	token := c.getToken()

	//读取存储好的音频文件
	voiceFile, err := os.Open("/root/go/src/talkGo/static/" + silkFileName + ".pcm")
	//voiceFile, err := os.Open("static/" + h.Filename)
	if err != nil {
		fmt.Println(err)
	}
	var b = make([]byte, 1024 * 1024)
	length, frerr := voiceFile.Read(b)
	if frerr != nil {
		fmt.Println(frerr)
	}

	//发起转换成文字请求
	var voiceJs VoiceJson
	voiceJs.Format = "pcm"
	voiceJs.Rate = 16000
	voiceJs.Channel = 1
	voiceJs.Cuid = "iamatest"
	voiceJs.Token = token
	voiceJs.Speech = base64.StdEncoding.EncodeToString(b[:length])
	voiceJs.Len = length
	req := httplib.Post("http://vop.baidu.com/server_api")
	req.Debug(true)
	req.Header("Content-Type","application/json")

	//查看 POST 的 JSON 内容
	//byts,_ := json.Marshal(voiceJson)
	//fmt.Println(string(byts))
	_, eor := req.JSONBody(voiceJs)
	if eor != nil {
		fmt.Println(eor)
	}
	jsonStr, err := req.String()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(jsonStr);
	var voiceResStruct VoiceResStruct
	err = json.Unmarshal([]byte(jsonStr), &voiceResStruct);
	if err != nil {
		fmt.Println(err)
	}

	if len(voiceResStruct.Result) > 0 {
		//Redis 记录识别的PCM音频缓存
		rc, err := lib.Dial()
		if err != nil {
			fmt.Println(err)
		}
		defer rc.Close()

		cacheKey := "user_talk_list_" + strconv.Itoa(1)
		_, err = redis.Int64(rc.Do("lpush", cacheKey, beego.AppConfig.String("rooturl") + "static/" + silkFileName + ".pcm"))
		if err != nil {
			fmt.Println(err)
		}

		//将我说的话放到返回数据中
		jsonMap := make(map[string]string)
		jsonMap["whatisay"] = voiceResStruct.Result[0]

		c.Data["json"] = jsonMap
	}

	c.ServeJSON()
}


// Login ...
// @Title Login
// @Description get Talk by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Talk
// @Failure 403 :id is empty
// @router /login [get]
func (c *TalkController) Login()  {
	code := c.GetString("code")
	userinfoJson := c.GetString("userinfo")
	fmt.Print("userinfo ===> ")
	fmt.Println(userinfoJson)

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
		fmt.Println(err)
	}
	_, err = rc.Do("set", sessionCacheKey, string(sessionData) , 300 * time.Second)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sessionCacheKey)

	//保存用户
	var u User;
	u.Openid = wxSession.Openid
	u.Username = userinfo.NickName
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
func (c *TalkController) CheckLogin()  {
	session_key := c.GetString("session_key")
	rc, err := lib.Dial()
	if err != nil {
		fmt.Println(err)
		c.Data["json"] = error(err)
		c.ServeJSON()
	}

	sv, err := redis.String(rc.Do("get", session_key));
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("session_key`s value is  ==>", sv)

	c.Data["json"] = map[string]string{"session_key" : string(sv)}
	c.ServeJSON()
}

// Post ...
// @Title Create
// @Description create Talk
// @Param	body		body 	models.Talk	true		"body for Talk content"
// @Success 201 {object} models.Talk
// @Failure 403 body is empty
// @router / [post]
func (c *TalkController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Talk by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Talk
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TalkController) Say() {
	msg := c.GetString("msg")

	jsonMap,err :=c._say(msg)
	if err != nil {
		fmt.Println(err);
	}

	c.Data["json"] = jsonMap;
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Talk
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Talk
// @Failure 403
// @router / [get]
func (c *TalkController) GetAll() {
	o := orm.NewOrm()
	var lists []orm.ParamsList
	count, err := o.QueryTable("msg").ValuesList(&lists)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(lists, count)

	c.Data["json"] = map[string] []orm.ParamsList {"items" : lists}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Talk
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Talk	true		"body for Talk content"
// @Success 200 {object} models.Talk
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TalkController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Talk
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TalkController) Delete() {

}

func (c * TalkController) error(err error) map[string]string {
	return map[string]string {"err" : err.Error()}
}

func (c *TalkController) getToken() string {
	var bdTken BdToken
	now := time.Now().Unix()

	fout, err := os.Open("./token.json")
	if err != nil {
		if os.IsNotExist(err) {
			fout,err = os.Create("./token.json")
			if err != nil {
				fmt.Println("file create fail")
			}
		}
	} else {
		tokenBytes := make([]byte, 1024)
		len, err := fout.Read(tokenBytes)
		if err != nil {
			fmt.Println(err.Error())
			return ""
		}

		if len > 0 {
			err = json.Unmarshal(tokenBytes[:len], &bdTken)
			if err != nil {
				fmt.Println(err.Error())
				return ""
			}
			if bdTken.Expires_in > now {
				fmt.Println("from file")
				return bdTken.Access_token
			}
		}
	}

	//访问接口最新TOKEN
	apikey := beego.AppConfig.String("BdApiKey")
	secretkey := beego.AppConfig.String("BdSecretKey")
	res := httplib.Get("https://openapi.baidu.com/oauth/2.0/token?grant_type=client_credentials&client_id=" + apikey + "&client_secret=" + secretkey)

	jsonStr, err := res.String()


	err = json.Unmarshal([]byte(jsonStr),&bdTken)
	if err != nil {
		fmt.Println(err)
	}
	bdTken.Expires_in += now

	_jsonByte, err := json.Marshal(bdTken)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	fout.Write(_jsonByte)

	return bdTken.Access_token
}

func (c *TalkController) saveMsg(msg string, replyContent string, mp3url string) int64 {
	//获取 REIDS 中的数据
	//记录识别的PCM音频缓存
	rc, err := lib.Dial()
	if err != nil {
		fmt.Println(err)
	}
	defer rc.Close()

	cacheKey := "user_talk_list_" + strconv.Itoa(1)
	reply, err := rc.Do("lpop", cacheKey)
	if err != nil {
		fmt.Println(err)
	}

	whatisayPcm, err := redis.String(reply, err)
	if err != nil {
		fmt.Println(err)
	}

	o := orm.NewOrm()
	dbMsg := Msg{Whatisay : msg,  WhatisayPcm : whatisayPcm, ReplyContent : replyContent, Mp3Url : mp3url, CreateTime : time.Now().Unix()}
	insertId, err := o.Insert(&dbMsg)
	if err != nil {
		fmt.Println(err.Error())
	}

	return insertId
}

func (c *TalkController) _say(msg string) (map[string]string, error) {
	//接口访问
	txUrl := beego.AppConfig.String("TLApi") + "?userid=1&key=" + beego.AppConfig.String("TLKey") + "&info=" + msg
	req := httplib.Get(txUrl)
	resJson, err := req.String()
	fmt.Println(resJson)

	if err != nil {
		return nil, err
	}

	var tl Tl
	err = json.Unmarshal([]byte(resJson), &tl)
	if err != nil {
		return nil, err
	}

	jsonMap := make(map[string]string)
	jsonMap["code"] = strconv.Itoa(tl.Code)
	jsonMap["msg"] = tl.Text

	//合成语音
	token := c.getToken()
	audioUrl := beego.AppConfig.String("BdText2AudioApi") + "?tex=" + tl.Text + "&lan=zh&ctp=1&cuid=123321&per=4&tok=" + token
	res := httplib.Get(audioUrl)

	resp,err := res.Response()
	if err != nil {
		return nil, err
	}

	if resp.Header.Get("Content-Type") == "audio/mp3" {
		//保存文件
		mp3_id := strconv.Itoa(rand.Int())
		res.ToFile("./static/" + mp3_id + ".mp3")
		jsonMap["mp3"] = beego.AppConfig.String("rooturl") + "mp3dir/" + mp3_id + ".mp3"
		jsonMap["mp3_id"] = mp3_id

		//记录用户的提交的内容
		go c.saveMsg(msg, tl.Text, jsonMap["mp3"])

	} else {
		_, err := res.String()
		if err != nil {
			fmt.Println(err)
		}
		//记录用户的提交的内容
		go c.saveMsg(msg, tl.Text, "")
	}
	return jsonMap, nil;
}
