package models

import (
	"fmt"
)

type Msg struct {
	Id int
	UserId int `orm:"index"`
	Whatisay string `orm:"index"`
	WhatisayPcm string
	ReplyContent string `orm:"type(text)"`
	Mp3Url string
	CreateTime int64 `orm:"auto_now_add;type(timestamp)"`
}

func init() {
	fmt.Println("msg model is init")
}
