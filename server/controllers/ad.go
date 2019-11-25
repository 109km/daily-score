package controllers

import (
	"server/models"
	"server/types"
	"strconv"
)

// Operations about Users
type AdController struct {
	BaseController
}

func (this *AdController) GetAll() {

	date := this.GetString("date")
	start, _ := strconv.Atoi(this.GetString("start"))
	limit, _ := strconv.Atoi(this.GetString("limit"))
	resStatus := GetResponseStatusByName(types.SUCCESS)
	resData := types.NewDataJSON()

	if limit == 0 {
		limit = 10
	}

	var ads []*models.Ad
	var total int64

	if date == "" {
		ads, total = models.GetAllAds(start, limit)
	} else {
		ads, total = models.GetAdsByDate(date, start, limit)
	}

	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE")
	this.Ctx.Output.Header("Access-Control-Allow-Headers", "x-requested-with,content-type")
	// this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	resData["total"] = total
	resData["list"] = ads
	// this.Data["json"] = ads
	this.ServeResponse(resStatus, resData)
}

func (this *AdController) Add() {
	var q1 []string
	q1 = this.GetStrings("q1")
	q2 := this.GetString("q2")
	q3 := this.GetString("q3")
	phone := this.GetString("phone")
	name := this.GetString("name")

	id, err := models.AddOneAd(q1, q2, q3, name, phone)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = map[string]int64{"id": id}
	}

	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE")
	this.Ctx.Output.Header("Access-Control-Allow-Headers", "x-requested-with,content-type")
	// this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")

	this.ServeJSON()
}
