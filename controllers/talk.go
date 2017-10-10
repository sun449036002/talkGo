package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/httplib"
	"encoding/json"
	"fmt"
	"strconv"
	"math/rand"
	"os"
	"time"
	. "talkGo/models"
)


//图灵
type Tl struct {
	Code int
	Text string
}

//百度Token
type BdToken struct {
	Access_token string
	Refresh_token string
	Session_key string
	Scope string
	Session_secret string
	Expires_in int64
}

type WxSession struct {
	Openid string
	Session_key string
	Unionid string
}

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
}

// Login ...
// @Title Login
// @Description get Talk by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Talk
// @Failure 403 :id is empty
// @router /login [get]
func (c *TalkController) Login()  {
	code := c.GetString("code");

	url := beego.AppConfig.String("wxApiUrl") + "sns/jscode2session?appid=" + beego.AppConfig.String("wxSmallAppId") + "&secret=" + beego.AppConfig.String("wxSmallSecret") + "&js_code=" + code + "&grant_type=authorization_code";
	fmt.Println(url);
	req := httplib.Get(url);
	sessionJson,err:= req.String();
	if err != nil {
		c.Data["json"] = error(err);
		c.ServeJSON();
	}

	var wxSession WxSession;
	err = json.Unmarshal([]byte(sessionJson), &wxSession);
	if err != nil {
		c.Data["json"] = error(err);
		c.ServeJSON();
	}

	fmt.Println(wxSession);
	redis,err := cache.NewCache("redis", `{"key":"talkRedis","conn":"127.0.0.1:6379","dbNum":"0","password":""}`)
	if err != nil {
		fmt.Println(err);
		c.Data["json"] = error(err);
		c.ServeJSON();
	}

	redis.Put("sessionKey", wxSession, 300 * time.Second)

	c.Data["json"] = map[string]string{"code" : code}
	c.ServeJSON();
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
	msg := c.GetString("msg");

	//接口访问
	txUrl := beego.AppConfig.String("TLApi") + "?userid=1&key=" + beego.AppConfig.String("TLKey") + "&info=" + msg;
	req := httplib.Get(txUrl)
	resJson, err := req.String();

	if err != nil {
		c.Data["json"] = error(err);
		c.ServeJSON();
	}

	//fmt.Println("转换前", resJson);
	var tl Tl;
	err = json.Unmarshal([]byte(resJson), &tl);
	if err != nil {
		fmt.Println(err);
	}

	jsonMap := make(map[string]string);
	jsonMap["code"] = strconv.Itoa(tl.Code);
	jsonMap["msg"] = tl.Text;

	//合成语音
	token := c.getToken();
	audioUrl := beego.AppConfig.String("BdText2AudioApi") + "?tex=" + tl.Text + "&lan=zh&ctp=1&cuid=123321&per=4&tok=" + token;
	res := httplib.Get(audioUrl);

	resp,err := res.Response();
	if err != nil {
		fmt.Println(err);
	}

	if resp.Header.Get("Content-Type") == "audio/mp3" {
		//保存文件
		mp3_id := strconv.Itoa(rand.Int());
		res.ToFile("./static/" + mp3_id + ".mp3");
		jsonMap["mp3"] = beego.AppConfig.String("rooturl") + "mp3dir/" + mp3_id + ".mp3";
		jsonMap["mp3_id"] = mp3_id;

		//记录用户的提交的内容
		go c.saveMsg(msg, tl.Text, jsonMap["mp3"]);

	} else {
		audioJson, err := res.String();
		if err != nil {
			fmt.Println(err);
		}
		//记录用户的提交的内容
		go c.saveMsg(msg, tl.Text, "");

		fmt.Println(audioJson);
	}
	fmt.Println(audioUrl);

	c.Data["json"] = jsonMap;
	c.ServeJSON();

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
	o := orm.NewOrm();
	var lists []orm.ParamsList;
	count, err := o.QueryTable("msg").ValuesList(&lists);
	if err != nil {
		fmt.Println(err.Error());
	}
	fmt.Println(lists, count);

	c.Data["json"] = map[string] []orm.ParamsList {"items" : lists};
	c.ServeJSON();
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
	var bdTken BdToken;
	now := time.Now().Unix();

	fout, err := os.Open("./token.json");
	if err != nil {
		if os.IsNotExist(err) {
			fout,err = os.Create("./token.json");
			if err != nil {
				fmt.Println("file create fail");
			}
		}
	} else {
		tokenBytes := make([]byte, 1024);
		len, err := fout.Read(tokenBytes);
		if err != nil {
			fmt.Println(err.Error());
			return "";
		}

		//fmt.Println("len > 0", len, len > 0);
		if len > 0 {
			//fmt.Println("%x", tokenBytes);
			//fmt.Println("%x", tokenBytes[:len]);
			err = json.Unmarshal(tokenBytes[:len], &bdTken);
			if err != nil {
				fmt.Println(err.Error());
				return "";
			}
			if bdTken.Expires_in > now {
				fmt.Println("from file");
				return bdTken.Access_token;
			}
		}
	}

	//访问接口最新TOKEN
	apikey := beego.AppConfig.String("BdApiKey");
	secretkey := beego.AppConfig.String("BdSecretKey");
	res := httplib.Get("https://openapi.baidu.com/oauth/2.0/token?grant_type=client_credentials&client_id=" + apikey + "&client_secret=" + secretkey)

	jsonStr, err := res.String();


	err = json.Unmarshal([]byte(jsonStr),&bdTken)
	if err != nil {
		fmt.Println(err);
	}
	bdTken.Expires_in += now;

	_jsonByte, err := json.Marshal(bdTken);
	if err != nil {
		fmt.Println(err.Error());
		return "";
	}

	fout.Write(_jsonByte);

	return bdTken.Access_token;
}

func (c *TalkController) saveMsg(msg string, replyContent string, mp3url string) {
	o := orm.NewOrm();
	dbMsg := Msg{Txt : msg, ReplyContent : replyContent, Mp3Url : mp3url, CreateTime : time.Now().Unix()}
	insertId, err := o.Insert(&dbMsg);
	if err != nil {
		fmt.Println(err.Error());
	}
	fmt.Println("insert id === >", insertId);
}
