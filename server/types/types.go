package types

type ResponseData struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    DataJSON `json:"data"`
}

type DataJSON map[string]interface{}
