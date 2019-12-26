package autotask

import (
	"server/utils"
)

func StartDailySentence() {

	weatherData := utils.GetWeather() // 获取今日天气信息
	// sentenceData := utils.GetDailySentence() // 获取每日一句

	// sentenceArr := utils.SplitSentence(sentenceData["translation"].(string))

	lowestTemp, _ := weatherData["tem2"].(string)
	highestTemp, _ := weatherData["tem1"].(string)
	currentTemp, _ := weatherData["tem"].(string)
	dateStr, _ := weatherData["date"].(string)
	wea, _ := weatherData["wea"].(string)

	tmpStr1 := lowestTemp + "~" + highestTemp + "°C" // 温度范围
	tmpStr2 := currentTemp + "°C，天气：" + wea          // 当前温度
	pm25Str := weatherData["air_pm25"].(string)      // pm2.5

	// sentenceString := ""

	// sentenceLen := len(sentenceArr)
	// for i := 0; i < sentenceLen; i++ {
	// 	sentenceString += "\"" + sentenceArr[i] + "\""
	// 	if i < sentenceLen-1 {
	// 		sentenceString += ","
	// 	}
	// }

	params := "{\"PhoneNumberSet\":[\"+8613621253007\",\"+8615701503610\"],\"TemplateID\":\"506849\",\"Sign\":\"shulanbaobei\",\"TemplateParamSet\":[\"" + dateStr + "\",\"" + tmpStr1 + "\",\"" + tmpStr2 + "\",\"" + pm25Str + "\"],\"SmsSdkAppid\":\"1400298212\"}"
	utils.SendSMS(params)
}
