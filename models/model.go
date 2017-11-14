package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type Model struct {
}

func init()  {
	fmt.Println("base model is init regiser models in this model")
	orm.RegisterModel(new(Msg), new(User));

	// 注册驱动
	orm.RegisterDataBase("default", "mysql", "root:mm@tcp(127.0.0.1:3306)/sun?charset=utf8");

	dropOldTable := false;
	orm.RunSyncdb("default", dropOldTable, true);
}
