package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"

)

type Msg struct {
	Id int
	Whatisay string `orm:"index"`
	WhatisayPcm string
	ReplyContent string `orm:"type(text)"`
	Mp3Url string
	CreateTime int64 `orm:"auto_now_add;type(timestamp)"`
	User *User `orm:"rel(fk)"`
}

func init() {
	fmt.Println("msg model is init")
}

func (m *Msg) GetMsgList() []Msg  {
	o := orm.NewOrm()
	var msgList []Msg
	o.QueryTable("msg").Limit(10, 0).RelatedSel().All(&msgList)
	fmt.Println(msgList)

	return msgList
}