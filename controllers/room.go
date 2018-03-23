package controllers

import (
	"talkGo/models"
	"strconv"
)

type RoomController struct {
	BaseController
}

// URLMapping ...
func (c *RoomController) URLMapping() {
	c.Mapping("Create", c.Create)
	c.Mapping("GetList", c.GetList)
	c.Mapping("IExit", c.IExit)
}

// 创建房间...
// @Title Create
// @Description up voice to server,chnage to text
// @Success 200 {object} models.Talk
// @Failure 403 :id is empty
// @router /Create [get]
func (c *RoomController) Create() {
	c.init()
	jsonMap := make(map[string]interface{})

	name := c.GetString("name")
	if name == "" {
		jsonMap["code"] = "100"
		jsonMap["msg"] = "empty name"
		c.Data["json"] = jsonMap
		c.ServeJSON()
		return
	}

	roomModel := new(models.Room)
	roomId, err := roomModel.Create(c.user.Id, name)
	if err != nil {
		jsonMap["code"] = "100"
		jsonMap["msg"] = "room create failed"
		c.Data["json"] = jsonMap
		c.ServeJSON()
		return
	}

	jsonMap["code"] = "0"
	jsonMap["msg"] = "success"
	jsonMap["roomId"] = "room_" + strconv.Itoa(int(roomId))
	c.Data["json"] = jsonMap
	c.ServeJSON()
}

// 房间列表...
// @Title GetList
// @Description up voice to server,chnage to text
// @Success 200 {object} models.Talk
// @Failure 403 :id is empty
// @router /get-list [get]
func (c *RoomController) GetList() {
	jsonMap := make(map[string]interface{})

	page, _ := c.GetInt64("page", 1)

	roomModel := new(models.Room)
	roomList, nextPage, isEnd := roomModel.GetList(page)

	jsonMap["code"] = "0"
	jsonMap["msg"] = "success"
	jsonMap["items"] = roomList
	jsonMap["page"] = nextPage
	jsonMap["isEnd"] = isEnd
	jsonMap["userId"] = c.user.Id
	c.Data["json"] = jsonMap
	c.ServeJSON()
}


// 房主离开...
// @Title IExit
// @Description up voice to server,chnage to text
// @Success 200 {object} models.Talk
// @Failure 403 :id is empty
// @router /i-exit [get]
func (c *RoomController) IExit() {
	jsonMap := make(map[string]interface{})

	roomId, _ := c.GetInt("roomId", 0)

	roomModel := new(models.Room)
	ok,_ := roomModel.Exit(c.user.Id, roomId)

	jsonMap["code"] = 0
	jsonMap["msg"] = "success"
	if !ok {
		jsonMap["code"] = 0
		jsonMap["msg"] = "exit room fail"
	}

	c.Data["json"] = jsonMap
	c.ServeJSON()
}