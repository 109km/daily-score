package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"server/types"
)

var WeatherAPI = "https://www.tianqiapi.com/api/"

type WeatherParams struct {
	AppId     string `json:"appid"`
	AppSecret string `json:"appsecret"`
	Version   string `json:"version"`
	CityId    string `json:"cityid"`
	City      string `json:"city"`
}

func GetWeather() types.DataJSON {
	query := url.Values{}
	query.Set("appid", "69761685")
	query.Set("appsecret", "qkmt79VE")
	query.Set("version", "v6")
	query.Set("cityid", "101010100")
	query.Set("city", "北京")

	url := WeatherAPI + "?" + query.Encode()
	data, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	weather, err2 := ioutil.ReadAll(data.Body)
	data.Body.Close()
	if err2 != nil {
		log.Fatal(err2)
	}
	var weatherJSON types.DataJSON
	err3 := json.Unmarshal([]byte(string(weather)), &weatherJSON)
	if err3 != nil {
		log.Fatal(err3)
	}
	return weatherJSON
}
