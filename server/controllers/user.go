package controllers

import (
	"fmt"
	"server/models"
	"strconv"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

func (u *UserController) Get() {
	uid, _ := strconv.Atoi(u.GetString(":uid"))
	if uid > 0 {
		user, err := models.GetUserById(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

func (this *UserController) Add() {

	mobile := this.GetString("mobile")
	password := this.GetString("password")
	nickname := this.GetString("nickname")

	if mobile == "" {
		response := NewResponseData(-1, "`mobile` must not be empty.")
		this.Data["json"] = response
		fmt.Println(response)
	} else {
		id, err := models.AddOneUser(mobile, password, nickname)
		if err != nil {
			this.Data["json"] = err.Error()
		} else {
			this.Data["json"] = map[string]int64{"uid": id}
		}
	}
	this.ServeJSON()
}
