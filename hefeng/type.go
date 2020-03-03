package hefeng

type WeatherInfoDetails struct {
	Brf string
	Txt string
}

type WeatherInfo struct {
	Comfort WeatherInfoDetails			//舒适度
	DressIndex WeatherInfoDetails		//穿衣指数
	ColdIndex WeatherInfoDetails		//感冒指数
	Sport WeatherInfoDetails			//运动指数
	Travel WeatherInfoDetails			//旅游指数
	Uv WeatherInfoDetails				//紫外线指数
	WashCarIndex WeatherInfoDetails		//洗车指数
	AirPollution WeatherInfoDetails		//大气污染
}

type HeFWeather struct {
	WeatherData WeatherInfo
}




