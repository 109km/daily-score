package types

var RESPONSE_STATUS = make(map[string]ResponseStatusStructure)

var SUCCESS = "SUCCESS"
var USERNAME_OR_PASSWORD_WRONG = "USERNAME_OR_PASSWORD_WRONG"
var CAN_NOT_GET_USERLIST = "CAN_NOT_GET_USERLIST"
var REQUEST_PARAMS_ERROR = "REQUEST_PARAMS_ERROR"

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
	RESPONSE_STATUS[REQUEST_PARAMS_ERROR] = ResponseStatusStructure{-10002, "传入参数错误"}

	// user
	RESPONSE_STATUS[USERNAME_OR_PASSWORD_WRONG] = ResponseStatusStructure{-50001, "用户名/密码错误"}
	RESPONSE_STATUS[CAN_NOT_GET_USERLIST] = ResponseStatusStructure{-50002, "无法获取用户列表"}
}
