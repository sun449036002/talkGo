package controllers

import (
	"talkGo/models"
	"strconv"
	"path"
	"github.com/astaxie/beego"
	"time"
)

type RoomController struct {
	BaseController
}

// URLMapping ...
func (c *RoomController) URLMapping() {
	c.Mapping("Create", c.Create)
	c.Mapping("GetList", c.GetList)
	c.Mapping("IExit", c.IExit)
	c.Mapping("UploadImg", c.UploadImg)
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
// @router /list [get]
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
// @router /iexit [get]
func (c *RoomController) IExit() {
	c.init()
	
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

// 图片上传...
// @Title UploadImg
// @Description up voice to server,chnage to text
// @router /uploadimg [post]
func (c *RoomController) UploadImg() {
	jsonMap := make(map[string]interface{})

	f, _, err := c.GetFile("cover")
	if err != nil {
		jsonMap["code"] = 0
		jsonMap["msg"] = "图片上传失败"
		c.Data["json"] = jsonMap
		c.ServeJSON()
		return
	}
	defer f.Close()

	filename := time.Now().String() + ".png"
	err = c.SaveToFile("file", path.Join("static/upload",filename))  //保存文件的路径。保存在static/upload中   （文件名）
	if err != nil {
		jsonMap["code"] = 0
		jsonMap["msg"] = "图片保存失败"
	} else {
		jsonMap["code"] = 0
		jsonMap["msg"] = "图片上传成功"
		jsonMap["path"] = beego.AppConfig.String("rooturl") + "static/upload/" + filename
	}

	c.Data["json"] = jsonMap
	c.ServeJSON()
}