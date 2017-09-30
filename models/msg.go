package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)


type Msg struct {
	Id int
	Txt string
	CreateTime int64
}

func init() {
	orm.RegisterModel(new(Msg));

	// 注册驱动
	orm.RegisterDataBase("default", "mysql", "root:mm@tcp(127.0.0.1:3306)/sun?charset=utf8");

	orm.RunSyncdb("default", false, true);
}
