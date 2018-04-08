package controllers

import "fmt"

type NotifyController struct {
	BaseController
}


// URLMapping ...
func (c *NotifyController) URLMapping() {
	c.Mapping("OnPublish", c.OnPublish)
	c.Mapping("OnPublishDone", c.OnPublishDone)
	c.Mapping("OnPlay", c.OnPlay)
	c.Mapping("OnPlayDone", c.OnPlayDone)
}

// 直播推流时通知...
// @Title OnPublish
// @Description up voice to server,chnage to text
// @Success 200 {object} models.Talk
// @Failure 403 :id is empty
// @router /onPublish [get]
func (c *NotifyController) OnPublish() {
	fmt.Println("OnPublish")
}


// 直播推流结束时通知...
// @Title OnPublishDone
// @Description up voice to server,chnage to text
// @Success 200 {object} models.Talk
// @Failure 403 :id is empty
// @router /onPublishDone [get]
func (c *NotifyController) OnPublishDone() {
	fmt.Println("OnPublishDone")
}


// 直播拉流时通知...
// @Title OnPlay
// @Description up voice to server,chnage to text
// @Success 200 {object} models.Talk
// @Failure 403 :id is empty
// @router /OnPlay [get]
func (c *NotifyController) OnPlay() {
	fmt.Println("OnPlay")
}

// 直播拉流结束时通知...
// @Title OnPlayDone
// @Description up voice to server,chnage to text
// @Success 200 {object} models.Talk
// @Failure 403 :id is empty
// @router /OnPlayDone [get]
func (c *NotifyController) OnPlayDone() {
	fmt.Println("OnPlayDone")
}
