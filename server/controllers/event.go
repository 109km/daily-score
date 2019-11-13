package controllers

import (
	"server/models"
	"server/types"
	"strconv"
)

// Operations about object
type EventController struct {
	BaseController
}

func (this *EventController) GetAll() {
	tasks := models.GetAllEvents()
	this.Data["json"] = tasks
	this.ServeJSON()
}

func (this *EventController) Get() {
	id, _ := strconv.ParseInt(this.GetString(":eid"), 10, 64)
	if id > 0 {
		event, err := models.GetEvent(id)
		if err != nil {
			this.Data["json"] = err.Error()
		} else {
			this.Data["json"] = event
		}
	}
	this.ServeJSON()
}

func (this *EventController) Add() {

	title := this.GetString("title")
	level := this.GetString("level")
	totalScore := this.GetString("total_score")
	startTime := this.GetString("start_time")
	endTime := this.GetString("end_time")

	userSession := this.GetSession(USER_COOKIE_SESSION_ID)

	if userSession == nil {
		resStatus := GetResponseStatusByName(types.USER_NOT_LOGIN)
		this.ServeResponse(resStatus, nil)
	}

	resStatus := GetResponseStatusByName(types.SUCCESS)
	resData := types.NewDataJSON()

	if title == "" ||
		level == "" ||
		totalScore == "" ||
		startTime == "" ||
		endTime == "" {
		res := GetResponseStatusByName(types.REQUEST_PARAMS_ERROR)
		res.Message = "Some parameters are missing."
		this.ServeResponse(res, nil)
	} else {
		totalScoreNum, _ := strconv.Atoi(totalScore)
		levelNum, _ := strconv.Atoi(level)
		id, err := models.AddOneEvent(title, levelNum, totalScoreNum, startTime, endTime)
		if err == nil {
			resData["id"] = id
		} else {
			resStatus = GetResponseStatusByName(types.EVENT_CREATED_ERROR)
		}
		this.ServeResponse(resStatus, resData)
	}
}
