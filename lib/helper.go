package lib

import (
	"time"
	"math/rand"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"fmt"
	"encoding/json"
	"talkGo/structs"
)

//生成随机字符串
func GetRandomString(_len int) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	length := len(bytes)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < _len; i++ {
		result = append(result, bytes[r.Intn(length)])
	}

	return string(result)
}

//连接 redis
func Dial() (redis.Conn, error){
	return redis.Dial("tcp", "127.0.0.1:6379")
}


//获取 access token
func GetAccessToken() structs.AccessToken {
	url := beego.AppConfig.String("wxApiUrl") + "cgi-bin/token?grant_type=client_credential&appid=" + beego.AppConfig.String("wxSmallAppId") + "&secret=" + beego.AppConfig.String("wxSmallSecret")
	req := httplib.Get(url)
	accessTokenJson,err:= req.String()
	fmt.Println("acessTokenJson =====> ", accessTokenJson)
	if err != nil {
		fmt.Println(err.Error())
	}

	var accessToken structs.AccessToken
	err = json.Unmarshal([]byte(accessTokenJson), &accessToken)
	if err != nil {
		fmt.Println(err.Error())
	}

	return accessToken
}