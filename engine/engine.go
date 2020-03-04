package engine

import (
	"fmt"
	"sync"
	"time"
	"weatherAssistant/fetcher"
)

type Info struct {
	WorkerCount int
	Inf         Inform
	UserList    []string
	GettingData GetData
}

type DataType struct {
	WeatherInfo interface{}
}

type GetData interface {
	Parse(contents []byte) bool
	DisplayData() string
}

type receiveInfo struct {
	data     string
	receiver string
}

func (e Info) worker(id int, jobs <-chan receiveInfo, wg *sync.WaitGroup) {

	for {
		for j := range jobs {
			e.Inf.Inform(j.data, j.receiver)
			wg.Done()
		}
	}
}

func (e Info) Run(seeds Request) {

	jobs := make(chan receiveInfo, 10)

	//获取天气数据
	body, err := fetcher.Fetch(seeds.Url)

	if err != nil {
		fmt.Println("fetcher data err")
		e.Inf.Inform("获取天气数据错误，请立刻查看错误", "337612001@qq.com")
		return
	} else {
		fmt.Println("fetcher data normal")
		//解析天气数据
		if e.GettingData.Parse(body) {
			fmt.Println("Parse data normal")
		} else {
			fmt.Println("Parse data err")
			e.Inf.Inform("解析天气数据错误，请立刻查看错误", "337612001@qq.com")
			return
		}
	}

	sendBuf := e.GettingData.DisplayData()

	var wg sync.WaitGroup
	//开启通知线程
	for w := 1; w <= e.WorkerCount; w++ {
		go e.worker(w, jobs, &wg)
	}

	time.Sleep(time.Second)

	//通知用户
	for j := 0; j < len(e.UserList); j++ {
		wg.Add(1)
		rr := receiveInfo{}

		rr.receiver = e.UserList[j]
		rr.data = sendBuf
		jobs <- rr
	}

	wg.Wait()

	close(jobs)
}
