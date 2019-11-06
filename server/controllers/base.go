package controllers

import (
	"server/types"

	"github.com/astaxie/beego"
)

var (
	RESPONSE_STATUS_MESSAGE map[int]string
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) ServeResponse(code int, message string, data types.DataJSON) {
	response := types.NewResponseData(code, message, data)
	this.Data["json"] = response
	this.ServeJSON()
	this.StopRun()
}

func GetResponseMessageByCode(code int) string {
	return RESPONSE_STATUS_MESSAGE[code]
}

func init() {
	RESPONSE_STATUS_MESSAGE = make(map[int]string)
	RESPONSE_STATUS_MESSAGE[0] = "success"
	RESPONSE_STATUS_MESSAGE[-50001] = "用户名/密码错误"
}
