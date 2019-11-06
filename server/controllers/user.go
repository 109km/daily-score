package controllers

import (
	"fmt"
	"server/models"
	"server/types"
	"strconv"
)

// Operations about Users
type UserController struct {
	BaseController
}

func (this *UserController) GetAll() {
	resStatus := GetResponseStatusByName(types.SUCCESS)
	resData := types.NewDataJSON()
	users, err := models.GetAllUsers()

	// sessionKey := this.Ctx.GetSecureCookie(beego.BConfig.WebConfig.Session.SessionName)
	userSession := this.GetSession(USER_COOKIE_SESSION_ID)

	fmt.Println(USER_COOKIE_SESSION_ID, userSession)

	if err != nil {
		resStatus = GetResponseStatusByName(types.CAN_NOT_GET_USERLIST)
	} else {
		resData["list"] = users
		resData["total"] = len(users)
		resData["cookie"] = USER_COOKIE_SESSION_ID

		if userSession != nil {
			resData["sessionId"] = userSession.(int)
		}

	}

	this.ServeResponse(resStatus, resData)
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
		res := GetResponseStatusByName(types.REQUEST_PARAMS_ERROR)
		res.Message = "`mobile` must not be empty."
		this.ServeResponse(res, nil)
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

	resStatus := GetResponseStatusByName(types.SUCCESS)
	resData := types.NewDataJSON()

	userSession := this.GetSession(USER_COOKIE_SESSION_ID)

	if userSession == nil {

		_, err := models.LoginUser(mobile, password)
		if err == nil {
			this.SetSession(USER_COOKIE_SESSION_ID, int(1))
			resData["mobile"] = mobile
			resData["sessionId"] = int(1)
		} else {
			resStatus = GetResponseStatusByName(types.USERNAME_OR_PASSWORD_WRONG)
		}
	} else {
		this.SetSession(USER_COOKIE_SESSION_ID, userSession.(int))
		resStatus.Message = "not expired"
		resData["mobile"] = mobile
		resData["sessionId"] = userSession.(int)
	}
	this.ServeResponse(resStatus, resData)

}
