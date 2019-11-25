package types

var RESPONSE_STATUS = make(map[string]ResponseStatusStructure)

var SUCCESS = "SUCCESS"

var REQUEST_PARAMS_ERROR = "REQUEST_PARAMS_ERROR"

var CAN_NOT_GET_USERLIST = "CAN_NOT_GET_USERLIST"
var USER_NOT_LOGIN = "USER_NOT_LOGIN"
var USERNAME_OR_PASSWORD_WRONG = "USERNAME_OR_PASSWORD_WRONG"

var EVENT_CREATED_ERROR = "EVENT_CREATED_ERROR"
var EVENTS_FOUND_ERROR = "EVENTS_FOUND_ERROR"

type ResponseStatusStructure struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseData struct {
	ResponseStatusStructure
	Data DataJSON `json:"data"`
}

type Any interface{}

type DataJSON map[string]Any

func NewDataJSON() DataJSON {
	return make(DataJSON)
}

func NewResponseData(rs ResponseStatusStructure, data DataJSON) *ResponseData {
	return &ResponseData{rs, data}
}

func init() {
	// success
	RESPONSE_STATUS[SUCCESS] = ResponseStatusStructure{0, "success"}

	// params
	RESPONSE_STATUS[REQUEST_PARAMS_ERROR] = ResponseStatusStructure{-1001, "传入参数错误"}

	// user
	RESPONSE_STATUS[USERNAME_OR_PASSWORD_WRONG] = ResponseStatusStructure{-2001, "用户名/密码错误"}
	RESPONSE_STATUS[USER_NOT_LOGIN] = ResponseStatusStructure{-2010, "用户未登录"}

	RESPONSE_STATUS[CAN_NOT_GET_USERLIST] = ResponseStatusStructure{-2100, "无法获取用户列表"}

	// event
	RESPONSE_STATUS[EVENT_CREATED_ERROR] = ResponseStatusStructure{-3001, "事件创建错误"}
	RESPONSE_STATUS[EVENTS_FOUND_ERROR] = ResponseStatusStructure{-3100, "事件列表查找错误"}

}
