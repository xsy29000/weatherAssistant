package main

import (
	"fmt"
	"gopkg.in/gcfg.v1"
	"time"
	"weatherAssistant/engine"
	"weatherAssistant/hefeng"
	"weatherAssistant/inform"
)

func main() {

	//定時任務
	//c := cron.New()
	//spec := "0 43 17 * * ?"
	//c.AddFunc(spec, func() {
	t1 := time.Now().Year()   //年
	t2 := time.Now().Month()  //月
	t3 := time.Now().Day()    //日
	t4 := time.Now().Hour()   //小时
	t5 := time.Now().Minute() //分钟
	fmt.Printf("%d年%d月%d日 %d:%d\n", t1, t2, t3, t4, t5)

	UserList := []string{"337612001@qq.com"}

	var weatherInfo hefeng.HeFWeather

	//读取邮件发送模块配置
	var mail inform.MailIniStr

	iniInfo := struct {
		HeFeng struct {
			Url string
		}
		Email struct {
			Sender         string
			SenderPassword string
		}
	}{}
	err := gcfg.ReadFileInto(&iniInfo, "config.ini")
	if err != nil {
		fmt.Printf("Failed to parse config file: %s\n", err)
		return
	}

	mail = iniInfo.Email
	e := engine.Info{
		WorkerCount: 3,
		Inf: inform.Mail{
			MailIni: mail,
		},
		UserList:    UserList,
		GettingData: &weatherInfo,
	}

	e.Run(engine.Request{
		Url: iniInfo.HeFeng.Url,
	})
	//})
	//
	//c.Start()
	//
	//select {}

	return
}
