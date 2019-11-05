package controllers

import (
	"server/models"
	"strconv"
)

// Operations about object
type TaskController struct {
	BaseController
}

func (this *TaskController) GetAll() {
	tasks := models.GetAllTasks()
	this.Data["json"] = tasks
	this.ServeJSON()
}

func (this *TaskController) Get() {
	id, _ := strconv.Atoi(this.GetString(":tid"))
	if id > 0 {
		task, err := models.GetTask(id)
		if err != nil {
			this.Data["json"] = err.Error()
		} else {
			this.Data["json"] = task
		}
	}
	this.ServeJSON()
}
