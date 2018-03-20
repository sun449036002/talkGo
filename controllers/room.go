package controllers

import "talkGo/models"

type RoomController struct {
	BaseController
}

// URLMapping ...
func (c *RoomController) URLMapping() {
	c.Mapping("Create", c.Create)
	c.Mapping("GetList", c.GetList)
}

// 创建房间...
// @Title Create
// @Description up voice to server,chnage to text
// @Success 200 {object} models.Talk
// @Failure 403 :id is empty
// @router /Create [get]
func (c *RoomController) Create() {
	//c.init()
	jsonMap := make(map[string]interface{})

	name := c.GetString("name")
	if name == "" {
		jsonMap["code"] = "100"
		jsonMap["msg"] = "empty name"
		c.Data["json"] = jsonMap
		c.ServeJSON()
	}

	roomModel := new(models.Room)
	roomId, err := roomModel.Create(1, name)
	if err != nil {
		jsonMap["code"] = "100"
		jsonMap["msg"] = "room create failed"
		c.Data["json"] = jsonMap
		c.ServeJSON()
	}

	jsonMap["code"] = "0"
	jsonMap["msg"] = "success"
	jsonMap["roomId"] = roomId
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

	roomModel := new(models.Room)
	roomList, page, isEnd := roomModel.GetList(1)

	jsonMap["code"] = "0"
	jsonMap["msg"] = "success"
	jsonMap["items"] = roomList
	jsonMap["page"] = page
	jsonMap["isEnd"] = isEnd
	c.Data["json"] = jsonMap
	c.ServeJSON()
}
