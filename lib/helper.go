package lib

import (
	"time"
	"math/rand"
	"github.com/garyburd/redigo/redis"
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

func Dial() (redis.Conn, error){
	return redis.Dial("tcp", "127.0.0.1:6379")
}