package main

import (
	_ "daily-score-server/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init(){
	// Register db
	orm.RegisterDriver("mysql", orm.DRMySQL)
  orm.RegisterDataBase("daily_score", "mysql", "root:localhost@/daily_score?charset=utf8")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	

	beego.Run()
}
