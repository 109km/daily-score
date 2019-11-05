package controllers

// Operations about Index
type IndexController struct {
	BaseController
}

// @Title GetAll
// @Description get index
// @Success 200 {object}
// @router / [get]
func (this *IndexController) GetAll() {
	data := make(map[string]interface{})
	data["name"] = "hello"
	this.ServeResponse(0, "success", data)
}
