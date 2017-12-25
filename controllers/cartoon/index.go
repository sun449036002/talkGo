package cartoon

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/json-iterator/go"
)

type History struct {
	Id string
	Title string
	Time string
	ThumbnailList []string
	Link string
}

type IndexController struct {
	beego.Controller
}

func (c *IndexController) URLMapping() {
	c.Mapping("getList", c.GetList)
}

// getList 获取漫画列表...
// @router /list [get]
func (c *IndexController) GetList() {
	fmt.Println("this is get list")
	c.Auth()
}

func (c *IndexController) Auth() {
	/**
	type=/category/weimanhua/kbmh 恐怖漫画
	type=/category/weimanhua/gushimanhua 故事漫画
	type=/category/duanzishou 段子手
	type=/category/lengzhishi 冷知识
	type=/category/qiqu 奇趣
	type=/category/dianying 电影
	type=/category/gaoxiao 搞笑
	type=/category/mengchong 萌宠
	type=/category/xinqi 新奇
	type=/category/huanqiu 环球
	type=/category/sheying 摄影
	type=/category/wanyi 玩艺
	type=/category/chahua 插画
	 */
	res := httplib.Get("http://route.showapi.com/958-1?type=/category/weimanhua/kbmh&showapi_appid=" + beego.AppConfig.String("YiYuanAppId") + "&showapi_sign=" + beego.AppConfig.String("YiYuanSecretKey") + "&page=1")
	//fmt.Println(res.String())

	bts, err := res.Bytes()
	if(err != nil) {
		fmt.Println(err)
	}

	//jsoniter.ParseString()
	contentList := jsoniter.Get(bts, "showapi_res_body", "pagebean", "contentlist").GetInterface()
	hasMorePage := jsoniter.Get(bts, "showapi_res_body", "pagebean", "hasMorePage").ToBool()

	c.Data["list"] = contentList
	c.Data["hasMorePage"] = hasMorePage

	c.ServeJSON()
}