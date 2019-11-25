package controllers

import (
	"server/types"

	"github.com/astaxie/beego"
)

var (
	RESPONSE_STATUS_MESSAGE map[int]string
	USER_COOKIE_SESSION_ID  string
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) ServeResponse(rs types.ResponseStatusStructure, data types.DataJSON) {
	// rs := types.ResponseStatusStructure{code, message}
	response := types.NewResponseData(rs, data)
	this.Data["json"] = response
	this.ServeJSON()
	this.StopRun()
}

func GetResponseStatusByName(name string) types.ResponseStatusStructure {
	return types.RESPONSE_STATUS[name]
}
func (this *BaseController) Prepare() {
	USER_COOKIE_SESSION_ID = this.Ctx.GetCookie(beego.BConfig.WebConfig.Session.SessionName)
}
