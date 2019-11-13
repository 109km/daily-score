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
			beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
		),
		beego.NSNamespace("/task",
			beego.NSRouter("/", &controllers.TaskController{}, "get:GetAll"),
			beego.NSRouter("/:tid", &controllers.TaskController{}, "get:Get"),
		),
		beego.NSNamespace("/event",
			beego.NSRouter("/", &controllers.EventController{}, "get:GetAll"),
			beego.NSRouter("/:eid", &controllers.EventController{}, "get:Get"),
			beego.NSRouter("/add", &controllers.EventController{}, "post:Add"),
		),
		beego.NSNamespace("/ad",
			beego.NSRouter("/", &controllers.AdController{}, "get:GetAll"),
			beego.NSRouter("/add", &controllers.AdController{}, "post:Add"),
		),
	)
	beego.AddNamespace(ns)
}
