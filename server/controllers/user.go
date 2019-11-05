package controllers

import (
	"fmt"
	"server/models"
	"strconv"
)

// Operations about Users
type UserController struct {
	BaseController
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
		response := NewResponseData(-1, "`mobile` must not be empty.", make(DataJSON))
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

func (this *UserController) Login() {

	username := this.GetString("username")
	// password := this.GetString("password")

	userSession := this.GetSession(username)

	if userSession == nil {
		this.SetSession(username, int(1))
		resCode := 0
		resMessage := "success"
		resData := make(DataJSON)
		resData["username"] = username
		resData["sessionId"] = int(1)
		this.ServeResponse(resCode, resMessage, resData)

	} else {
		this.SetSession("asta", userSession.(int)+1)
		this.Data["num"] = userSession.(int)
	}

	this.ServeJSON()
}
