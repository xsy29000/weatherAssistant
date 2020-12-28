package hefeng

type Display struct{}

func (p HeFWeather) DisplayData() string {
	var data string

	data += stringJoint("以下是为您推送今日的生活指数信息：", "")
	data += stringJoint("舒适度: ", p.WeatherData.Comfort.Txt)
	data += stringJoint("穿衣指数: ", p.WeatherData.DressIndex.Txt)
	data += stringJoint("感冒指数: ", p.WeatherData.ColdIndex.Txt)
	data += stringJoint("运动指数: ", p.WeatherData.Sport.Txt)
	data += stringJoint("旅游指数: ", p.WeatherData.Travel.Txt)
	data += stringJoint("紫外线指数: ", p.WeatherData.Uv.Txt)
	data += stringJoint("洗车指数: ", p.WeatherData.WashCarIndex.Txt)
	data += stringJoint("大气污染: ", p.WeatherData.AirPollution.Txt)

	return data
}

func stringJoint(firstData string, secondData string) string {
	var result string

	result += "<h3>"
	result += firstData
	result += secondData
	result += "</h3>"

	return result
}
