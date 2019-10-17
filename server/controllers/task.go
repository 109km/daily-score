package controllers

import (
	"server/models"

	"github.com/astaxie/beego"
)

// Operations about object
type TaskController struct {
	beego.Controller
}
// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (this *TaskController) GetAll() {
	tasks := models.GetAllTasks()
	this.Data["json"] = tasks
	this.ServeJSON()
}
