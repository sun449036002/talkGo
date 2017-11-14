package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Msg struct {
	Id int
	Txt string `orm:"index"`
	ReplyContent string `orm:"type(text)"`
	Mp3Url string
	CreateTime int64 `orm:"auto_now_add;type(timestamp)"`
}

func init() {
	orm.RegisterModel(new(Msg), new(User));

	// 注册驱动
	orm.RegisterDataBase("default", "mysql", "root:mm@tcp(127.0.0.1:3306)/sun?charset=utf8");

	dropOldTable := false;
	orm.RunSyncdb("default", dropOldTable, true);
}
