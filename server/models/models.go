package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	OrmInstance orm.Ormer
)

type User struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type Task struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	UnitScore  int    `json:"unit_score"`
	TotalScore int    `json:"total_score"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
}

func init() {
	// Register db
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:s09070825!@tcp(127.0.0.1:3306)/daily_score?charset=utf8")

	// 注册model
	orm.RegisterModel(new(User), new(Task))

	OrmInstance = orm.NewOrm()
}

func ProceedSearchError(resErr error) (err error) {
	if resErr == orm.ErrNoRows {
		return errors.New("查询不到")
	} else if resErr == orm.ErrMissPK {
		return errors.New("找不到主键")
	} else {
		return nil
	}
}
