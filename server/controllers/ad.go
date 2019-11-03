package controllers

import (
	"server/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type AdController struct {
	beego.Controller
}

func (this *AdController) GetAll() {

	date := this.GetString("date")

	var ads []*models.Ad

	if date == "" {
		ads = models.GetAllAds()
	} else {
		ads = models.GetAdsByDate(date)
	}

	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE")
	this.Ctx.Output.Header("Access-Control-Allow-Headers", "x-requested-with,content-type")
	// this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")

	this.Data["json"] = ads
	this.ServeJSON()
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
