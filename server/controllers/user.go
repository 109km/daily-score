package controllers

import (
	"server/models"
	"server/types"
	"strconv"
)

// Operations about Users
type UserController struct {
	BaseController
}

type jsonUser struct {
	data map[string]interface{}
}

func (this *UserController) GetAll() {
	resCode := 0
	resMessage := GetResponseMessageByCode(resCode)
	resData := types.NewDataJSON()
	users := models.GetAllUsers()
	resData["list"] = users
	resData["total"] = len(users)
	this.ServeResponse(resCode, resMessage, resData)
}

func (this *UserController) Get() {
	uid, _ := strconv.ParseInt(this.GetString(":uid"), 10, 64)
	if uid > 0 {
		user, err := models.GetUserById(uid)
		if err != nil {
			this.Data["json"] = err.Error()
		} else {
			this.Data["json"] = user
		}
	}
	this.ServeJSON()
}

func (this *UserController) Add() {

	mobile := this.GetString("mobile")
	password := this.GetString("password")
	nickname := this.GetString("nickname")

	if mobile == "" {
		this.ServeResponse(-1, "`mobile` must not be empty.", nil)
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

	mobile := this.GetString("mobile")
	password := this.GetString("password")

	userSession := this.GetSession(mobile)

	resCode := 0
	resMessage := GetResponseMessageByCode(resCode)
	resData := types.NewDataJSON()

	if userSession == nil {

		_, err := models.LoginUser(mobile, password)
		if err == nil {
			this.SetSession(mobile, int(1))
			resData["mobile"] = mobile
			resData["sessionId"] = int(1)
		} else {
			resCode = -50001
			resMessage = GetResponseMessageByCode(resCode)
		}
	} else {
		this.SetSession(mobile, userSession.(int))
		resMessage = "not expired"
		resData["mobile"] = mobile
		resData["sessionId"] = userSession.(int)
	}
	this.ServeResponse(resCode, resMessage, resData)

}
