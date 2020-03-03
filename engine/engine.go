package engine

import (
	"fmt"
	"sync"
	"time"
	"weatherAssistant/fetcher"
)

type Info struct {
	WorkerCount int
	Inf Inform
	UserList []string
	GettingData GetData
}

type DataType struct {
	WeatherInfo interface{}
}

type GetData interface {
	Parse(contents []byte)
	DisplayData()string
}

type receiveInfo struct {
	data string
	receiver string
}

func (e Info)worker(id int, jobs <-chan receiveInfo,wg *sync.WaitGroup){

	for {
		for j := range jobs {
			fmt.Println("worker ", id, "processing job ", j)
			time.Sleep(time.Second)
			e.Inf.Inform(j.data, j.receiver)
			wg.Done()
		}
	}
}


func(e Info)Run(seeds Request) {

	jobs := make(chan receiveInfo, 100)

	var wg sync.WaitGroup
	//开启通知线程
	for w := 1; w <= e.WorkerCount; w++ {
		go e.worker(w, jobs, &wg)
	}

	body, err := fetcher.Fetch(seeds.Url)

	//var result= WeatherInfo{}
	if err != nil {
		fmt.Println("err")
	} else {
		fmt.Println("no err")
		e.GettingData.Parse(body)
	}

	sendBuf := e.GettingData.DisplayData()

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