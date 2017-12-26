package cartoon

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/json-iterator/go"
)

type Cate struct {
	Title string
	Value string
}

type IndexController struct {
	beego.Controller
}

func (c *IndexController) URLMapping() {
	c.Mapping("getList", c.GetList)
	c.Mapping("GetDetail", c.GetDetail)
	c.Mapping("GetCategoryList", c.GetCategoryList)
}

// GetCategoryList 获取漫画列表...
// @router /cate-list [get]
func (c *IndexController) GetCategoryList()  {
	var cateList []Cate
	var cate1,cate2,cate3,cate4,cate5,cate6,cate7,cate8,cate9 Cate
	cate1.Title = "恐怖漫画"
	cate1.Value = "/category/weimanhua/kbmh"

	cate2.Title = "故事漫画"
	cate2.Value = "/category/weimanhua/gushimanhua"

	cate3.Title = "段子手"
	cate3.Value = "/category/duanzishou"

	cate4.Title = "奇趣"
	cate4.Value = "/category/qiqu"

	cate5.Title = "搞笑"
	cate5.Value = "/category/gaoxiao"

	//cate6.Title = "插画"
	//cate6.Value = "/category/chahua"
	//
	//cate7.Title = "新奇"
	//cate7.Value = "/category/xinqi"
	//
	//cate8.Title = "萌宠"
	//cate8.Value = "/category/mengchong"
	//
	//cate9.Title = "电影"
	//cate9.Value = "/category/dianying"

	cateList = append(cateList, cate1,cate2,cate3,cate4,cate5,cate6,cate7,cate8,cate9)

	c.Data["json"] = cateList
	c.ServeJSON()
}

// getList 获取漫画列表...
// @router /list [get]
func (c *IndexController) GetList() {
	fmt.Println("this is get list")
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
	page := c.GetString("page", "1")
	res := httplib.Get("http://route.showapi.com/958-1?type=/category/weimanhua/kbmh&showapi_appid=" + beego.AppConfig.String("YiYuanAppId") + "&showapi_sign=" + beego.AppConfig.String("YiYuanSecretKey") + "&page=" + page)
	//fmt.Println(res.String())

	bts, err := res.Bytes()
	if(err != nil) {
		fmt.Println(err)
	}

	contentList := jsoniter.Get(bts, "showapi_res_body", "pagebean", "contentlist").GetInterface()
	hasMorePage := jsoniter.Get(bts, "showapi_res_body", "pagebean", "hasMorePage").ToBool()

	json := make(map[string]interface{})
	json["list"] = contentList
	json["hasMorePage"] = hasMorePage
	c.Data["json"] = json

	c.ServeJSON()
}

// getList 获取漫画内容...
// @router /detail [get]
func (c *IndexController) GetDetail() {
	id := c.GetString("id")
	res := httplib.Get("http://route.showapi.com/958-2?showapi_appid=" + beego.AppConfig.String("YiYuanAppId") + "&showapi_sign=" + beego.AppConfig.String("YiYuanSecretKey") + "&id=" + id)

	bts, err := res.Bytes()
	if(err != nil) {
		fmt.Println(err)
	}

	content := jsoniter.Get(bts, "showapi_res_body", "item").GetInterface()
	json := make(map[string]interface{})
	json["content"] = content
	c.Data["json"] = json

	c.ServeJSON()
}