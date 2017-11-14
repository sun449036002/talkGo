package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"errors"
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
	if err == orm.ErrNoRows {
		id, err := o.Insert(user)
		if err != nil {
			fmt.Println("insert bad")
			return err;
		} else {
			fmt.Println("insert ok", id)
			return  nil;
		}
	}

	fmt.Println(u);
	return nil;
}