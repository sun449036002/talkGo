package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"github.com/astaxie/beego"
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
	num, _ := o.QueryTable("room").Filter("status", 1).Limit(pageSize, (page - 1) * pageSize).RelatedSel().OrderBy("-id").All(&roomList)

	list := make([]map[string]interface{}, len(roomList))
	for k,room := range roomList {
		tplMap := make(map[string]interface{})
		tplMap["roomId"] = "room_"  + strconv.Itoa(room.Id)
		tplMap["roomName"] = room.Name
		tplMap["roomCover"] = beego.AppConfig.String("rooturl") + "mp3dir/hlsCover/" + "room_"  + strconv.Itoa(room.Id) + "_cover.png"
		list[k] = tplMap
	}

	page++
	return list, page, num < pageSize
}

//创建房间
func (m *Room) Create(userId int, name string) (int64, error) {
	o := orm.NewOrm()
	m.Name = name
	m.UserId = userId
	m.Status = 1

	id, err := o.Insert(m)
	if err != nil {
		return 0, err
	}

	fmt.Println("the new room id is :", id)

	return id, nil
}

//更新房间状态
func (m *Room) Exit(userId int, roomId int) (bool, error) {
	o := orm.NewOrm()
	m.Id = roomId
	m.UserId = userId
	m.Status = 0
	num, err := o.Update(m, "status")
	if err != nil {
		return false, err
	}

	fmt.Println("the room`s status is :", m.Status, ",update rows = ", num)

	return true, nil
}