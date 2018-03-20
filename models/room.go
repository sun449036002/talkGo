package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Room struct {
	Id       int `orm:"auto;pk"`
	UserId  int
	Name string
	Status int
}


func init() {
	//RegisterModel 放在 msg Model中
	fmt.Println("room model init")
}

//房间列表
func (m *Room) GetList(page int64) ([]map[string]interface{}, int64, bool)  {
	var pageSize int64 = 10
	o := orm.NewOrm()
	var roomList []Room
	num, _ := o.QueryTable("room").Limit(pageSize, (page - 1) * pageSize).RelatedSel().OrderBy("-id").All(&roomList)

	list := make([]map[string]interface{}, 0)
	for k,room := range roomList {
		list[k]["roomId"] = "r_"  + strconv.Itoa(room.Id)
		list[k]["roomName"] = room.Name
	}

	return list, page, num < pageSize
}

//创建房间
func (m *Room) Create(userId int, name string) error {
	o := orm.NewOrm()
	m.Name = name
	m.UserId = userId
	id, err := o.Insert(m)
	if err != nil {
		return err
	}

	fmt.Println("the new room id is :", id)

	return nil
}