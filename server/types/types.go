package types

type ResponseData struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    DataJSON `json:"data"`
}

type Any interface{}

type DataJSON map[string]Any

func NewDataJSON() DataJSON {
	return make(DataJSON)
}

func NewResponseData(code int, message string, data DataJSON) *ResponseData {
	return &ResponseData{Code: code, Message: message, Data: data}
}
