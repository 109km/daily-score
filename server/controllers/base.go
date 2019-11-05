package controllers

import "github.com/astaxie/beego"

type ResponseData struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    DataJSON `json:"data"`
}

type DataJSON map[string]interface{}

func NewResponseData(code int, message string, data DataJSON) *ResponseData {
	return &ResponseData{Code: code, Message: message, Data: data}
}

type BaseController struct {
	beego.Controller
}

func (this *BaseController) ServeResponse(code int, message string, data DataJSON) {
	response := NewResponseData(code, message, data)
	this.Data["json"] = response
	this.ServeJSON()
	this.StopRun()
}
