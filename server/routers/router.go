package routers

import (
	"server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/index",
			beego.NSRouter("/", &controllers.IndexController{}, "get:GetAll"),
		),
		beego.NSNamespace("/user",
			beego.NSRouter("/", &controllers.UserController{}, "get:GetAll"),
			beego.NSRouter("/:uid", &controllers.UserController{}, "get:Get"),
			beego.NSRouter("/add", &controllers.UserController{}, "post:Add"),
		),
		beego.NSNamespace("/task",
			beego.NSRouter("/", &controllers.TaskController{}, "get:GetAll"),
			beego.NSRouter("/:tid", &controllers.TaskController{}, "get:Get"),
		),
	)
	beego.AddNamespace(ns)
}
