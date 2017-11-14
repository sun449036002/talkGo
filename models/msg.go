package models

import (
	"fmt"
)

type Msg struct {
	Id int
	Txt string `orm:"index"`
	ReplyContent string `orm:"type(text)"`
	Mp3Url string
	CreateTime int64 `orm:"auto_now_add;type(timestamp)"`
}

func init() {
	fmt.Println("msg model is init")
}
