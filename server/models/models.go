package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	OrmInstance orm.Ormer
)

type User struct {
	Id       int
	Nickname string
	Mobile   string
	Password string
}

type Task struct {
	Id         int
	Title      string
	UnitScore  int
	TotalScore int
	StartTime  string
	EndTime    string
}

func init() {
	// Register db
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:s09070825!@tcp(127.0.0.1:3306)/daily_score?charset=utf8")

	// 注册model
	orm.RegisterModel(new(User), new(Task))

	OrmInstance = orm.NewOrm()
}
