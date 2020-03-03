package hefeng

type Display struct{


}

func (d HeFWeather)DisplayData()string{
	var data string

	data += stringJoint("以下是为您推送今日的生活指数信息：","")


	data += stringJoint("舒适度: ",d.WeatherData.Comfort.Txt)
	data += stringJoint("穿衣指数: ",d.WeatherData.DressIndex.Txt)
	data += stringJoint("感冒指数: ",d.WeatherData.ColdIndex.Txt)
	data += stringJoint("旅游指数: ",d.WeatherData.Sport.Txt)
	data += stringJoint("旅游指数: ",d.WeatherData.Travel.Txt)
	data += stringJoint("紫外线指数: ",d.WeatherData.Uv.Txt)
	data += stringJoint("洗车指数: ",d.WeatherData.WashCarIndex.Txt)
	data += stringJoint("大气污染: ",d.WeatherData.AirPollution.Txt)

	return data
}




func stringJoint(firstData string,secondData string)string{
	var result string

	result += "<h3>"
	result += firstData
	result += secondData
	result += "</h3>"

	return result
}


