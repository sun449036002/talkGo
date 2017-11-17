package models

import (
	"fmt"
	"encoding/json"
	. "talkGo/structs"
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

func (m *Msg) GetMsgList(page int64) ([]Msg, int64, bool)  {
	var pageSize int64 = 10
	o := orm.NewOrm()
	var msgList []Msg
	num, _ := o.QueryTable("msg").Limit(pageSize, (page - 1) * pageSize).RelatedSel().All(&msgList)

	for _, msg := range msgList {
	    var userinfo Userinfo
	    err := json.Unmarshal([]byte(msg.User.UserJson), &userinfo)
	    if err != nil {
	        fmt.Println(err)
	    }
	    msg.User.Userinfo = userinfo
	}

	page ++
	return msgList, page, num < 10
}