package main

import (
	"server/autotask"
	_ "server/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
)

func main() {
	// fmt.Println(beego.BConfig.RunMode)
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
		// beego.BConfig.WebConfig.DirectoryIndex = true
		// beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// beego.BConfig.Listen.EnableAdmin = true

	// beego.BConfig.WebConfig.Session.SessionOn = true
	// beego.BConfig.WebConfig.Session.SessionName = "daily_score_session_id"
	// beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600

	// // Set session provider config
	// beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	// beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"

	go autotask.Start()

	beego.SetStaticPath("/public", "public")
	beego.Run()

}
