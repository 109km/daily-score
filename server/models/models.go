package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int
	Nickname string
	Mobile   string
	Password string
}

func init() {
	// Register db
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:s09070825!@tcp(127.0.0.1:3306)/daily_score?charset=utf8")

	// 注册model
	orm.RegisterModel(new(User))

	print("init in models.go")
}
