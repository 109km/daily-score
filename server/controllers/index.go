package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Users
type IndexController struct {
	beego.Controller
}

// @Title GetAll
// @Description get index
// @Success 200 {object} 
// @router / [get]
func (this *IndexController) GetAll() {
	this.Data["json"] = map[string]string {"Hello":"world"}
	this.ServeJSON()
}


