package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int
	Username string
	Openid string
	Profile  Profile
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

func (user *User) NewUser() bool  {
	u := User{Openid:user.Openid}

	o := orm.NewOrm()
	if o == nil {
		return false;
	}

	err := o.Read(&u)
	if err == orm.ErrNoRows {
		id, err := o.Insert(&user)
		if err != nil {
			fmt.Println("insert bad")
			return false;
		} else {
			fmt.Println("insert ok", id)
			return  true;
		}
	}

	fmt.Println(u);
	return true;
}