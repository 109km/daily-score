package controllers

import (
	types "server/types"
	"server/utils"
)

// Operations about Index
type IndexController struct {
	BaseController
}

func (this *IndexController) Index() {
	this.Ctx.WriteString("树懒宝贝欢迎你！")
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

func (this *IndexController) DailySMS() {

	weatherData := utils.GetWeather()        // 获取今日天气信息
	sentenceData := utils.GetDailySentence() // 获取每日一句

	sentenceArr := utils.SplitSentence(sentenceData["translation"].(string))
	sentenceString := ""
	sentenceLen := len(sentenceArr)
	for i := 0; i < sentenceLen; i++ {
		sentenceString += "\"" + sentenceArr[i] + "\""
		if i < sentenceLen-1 {
			sentenceString += ","
		}
	}
	weatherData["senctence"] = sentenceString
	res := GetResponseStatusByName(types.SUCCESS)
	this.ServeResponse(res, weatherData)
}
