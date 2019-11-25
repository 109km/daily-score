package controllers

import types "server/types"

// Operations about Index
type IndexController struct {
	BaseController
}

// @Title GetAll
// @Description get index
// @Success 200 {object}
// @router / [get]
func (this *IndexController) GetAll() {
	data := types.NewDataJSON()
	data["name"] = "hello"
	res := GetResponseStatusByName(types.SUCCESS)
	this.ServeResponse(res, data)
}
