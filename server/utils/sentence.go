package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"server/types"
)

var DailySentenseAPI = "https://apiv3.shanbay.com/weapps/dailyquote/quote/"

func GetDailySentence() types.DataJSON {
	var finalJSON types.DataJSON
	data, err := http.Get(DailySentenseAPI)
	if err != nil {
		log.Fatal(err)
	}
	datastring, err2 := ioutil.ReadAll(data.Body)
	data.Body.Close()
	if err2 != nil {
		log.Fatal(err2)
	}
	err3 := json.Unmarshal([]byte(string(datastring)), &finalJSON)
	if err3 != nil {
		log.Fatal(err3)
	}
	return finalJSON
}

func SplitSentence(sentence string) []string {
	// sentenceRune := []rune(sentence)
	// sentenceLen := len(sentenceRune)
	res := make([]string, 4)
	// s := ""
	// j := 0

	res[0] = "今天"
	res[1] = "你想"
	res[2] = "怎么"
	res[3] = "活？"

	// for i := 0; i < 48; i++ {
	// 	if i == sentenceLen {
	// 		break
	// 	}
	// 	fmt.Println(i % 12)
	// 	if i%12 == 11 || i == sentenceLen-1 {
	// 		res[j] = s
	// 		s = ""
	// 		j = j + 1
	// 	} else {
	// 		s = s + string(sentenceRune[i])
	// 	}
	// }
	return res
}
