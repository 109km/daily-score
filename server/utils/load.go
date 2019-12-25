package utils

import (
	"encoding/json"
	"io/ioutil"
)

func LoadJSON(filepath string) SecretConfig {
	file, _ := ioutil.ReadFile(filepath)
	var data SecretConfig
	err := json.Unmarshal(file, &data)
	if err != nil {
		panic(err.Error())
	}
	return data
}
