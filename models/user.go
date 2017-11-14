package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"talkGo/lib"
	"errors"
)

type User struct {
	Id       int
	Uri string `orm:"size(16);index;"`
	Username string
	Openid string
}

func init()  {
	fmt.Println("user model init")
}

func (user *User) NewUser() error  {
	u := User{Openid:user.Openid}

	fmt.Println(user.Openid)

	o := orm.NewOrm()
	if o == nil {
		fmt.Println("orm init faild inside");
		return errors.New("orm init faild");
	}

	err := o.Read(&u)
	fmt.Println(u)
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