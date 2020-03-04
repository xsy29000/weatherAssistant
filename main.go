package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
	"weatherAssistant/engine"
	"weatherAssistant/hefeng"
	"weatherAssistant/inform"
)

func main() {

	c := cron.New()
	spec := "0 43 17 * * ?"
	c.AddFunc(spec, func() {
		t1 := time.Now().Year()   //年
		t2 := time.Now().Month()  //月
		t3 := time.Now().Day()    //日
		t4 := time.Now().Hour()   //小时
		t5 := time.Now().Minute() //分钟
		fmt.Printf("%d年%d月%d日 %d:%d\n", t1, t2, t3, t4, t5)

		UserList := []string{"337612001@qq.com"}

		var weatherInfo hefeng.HeFWeather
		e := engine.Info{
			WorkerCount: 3,
			Inf:         inform.Mail{},
			UserList:    UserList,
			GettingData: &weatherInfo,
		}

		e.Run(engine.Request{
			Url: "https://free-api.heweather.net/s6/weather/lifestyle?location=guangzhou&key=26e6e05426344f8da738eb82bd30e574",
		})
	})

	c.Start()

	select {}

	return
}
