package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"talkGo/lib"
	"errors"
)

type User struct {
	Id       int `orm:"auto;pk"`
	Uri string `orm:"size(16);index;"`
	Username string
	Openid string
}

func init()  {
	//RegisterModel 放在 msg Model中
	fmt.Println("user model init")
}

func (user *User) NewUser() error  {
	u := new(User)
	u.Openid = user.Openid
	o := orm.NewOrm()
	if o == nil {
		fmt.Println("orm init faild inside");
		return errors.New("orm init faild");
	}

	//根据 openid 查询 ，默认是主键查询
	err := o.Read(u, "openid")
	fmt.Println("read error", err)
	if err == orm.ErrNoRows {
		user.Uri = lib.GetRandomString(16)
		id, err := o.Insert(user)
		if err != nil {
			fmt.Println("insert bad")
			return err;
		} else {
			fmt.Println("insert ok", id)
			return  nil;
		}
	}
	return nil;
}