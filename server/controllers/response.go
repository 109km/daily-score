package controllers

type ResponseData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewResponseData(code int, message string) *ResponseData {
	return &ResponseData{Code: code, Message: message}
}
