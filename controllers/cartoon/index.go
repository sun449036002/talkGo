package cartoon

import (
	"github.com/astaxie/beego"
	"fmt"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) URLMapping() {

}

// getList 获取漫画列表...
// @router /getList [get]
func getList() {
	fmt.Println("this is get list")
}