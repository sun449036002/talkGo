package controllers

import "talkGo/models"

type RoomController struct {
	BaseController
}

// URLMapping ...
func (c *RoomController) URLMapping() {
	c.Mapping("Create", c.Create)
	c.Mapping("GetList", c.Create)
}


//创建房间
func (c *RoomController) Create() {
	jsonMap := make(map[string]string)

	name := c.GetString("name")
	if name == "" {
		jsonMap["code"] = "100"
		jsonMap["msg"] = "empty name"
		c.Data["json"] = jsonMap
		c.ServeJSON()
	}



	jsonMap["code"] = "0"
	jsonMap["msg"] = "success"
	c.Data["json"] = jsonMap
	c.ServeJSON()
}

//房间列表
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
