package controllers

import (
	"server/types"

	"github.com/astaxie/beego"
)

func NewResponseData(code int, message string, data types.DataJSON) *types.ResponseData {
	return &types.ResponseData{Code: code, Message: message, Data: data}
}

type BaseController struct {
	beego.Controller
}

func (this *BaseController) ServeResponse(code int, message string, data types.DataJSON) {
	response := NewResponseData(code, message, data)
	this.Data["json"] = response
	this.ServeJSON()
	this.StopRun()
}
