package controllers

import "time"

/**
谜语
 */
type RiddleController struct {
	BaseController
}

func (c *RiddleController) URLMapping() {
	c.Mapping("Get", c.Get)
}

//谜语获取
// @Title Get
// @router /get [get]
func (c *RiddleController) Get() {
	jsonMap := make(map[string]interface{})

	now := time.Now()
	m := now.Minute()
	liveTimers := 600 - (m % 10)*60 - now.Second()

	jsonMap["code"] = 0
	jsonMap["liveTimers"] = liveTimers

	c.Data["json"] = jsonMap
	c.ServeJSON()
}
