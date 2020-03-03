package hefeng

import (
	"encoding/json"
	"fmt"
)

func (p *HeFWeather) Parse(contents []byte) {
	var returnData WeatherInfo

	fmt.Println("parse -------- >>>>>>>>  ", string(contents))

	m := make(map[string]interface{})
	json.Unmarshal(contents, &m)

	for key, value := range m {
		switch data := value.(type) {
		case []interface{}:
			fmt.Printf("map[%s]的值类型为[]interface, value = %T\n", key, data)
			for _, u := range data {
				fmt.Printf("map[%T]的值类型\n", u)
				m1 := make(map[string]interface{})
				m1 = u.(map[string]interface{})

				for _, value := range m1 {
					switch value.(type) {
					case []interface{}:
						for key, value = range m1 {
							//fmt.Printf("map2[%s %T]的值类型\n", key, value)
							switch data2 := value.(type) {
							case []interface{}:
								fmt.Printf("value2 = %v\n", data2)
								for _, j := range data2 {
									//fmt.Printf("mmll %T %T\n",i,j)
									m2 := make(map[string]interface{})
									m2 = j.(map[string]interface{})
									if vvv,ok := m2["type"]; vvv == "comf" &&ok{
										fmt.Println(m2["txt"])
										returnData.Comfort.Txt = m2["txt"].(string)
										returnData.Comfort.Brf = m2["brf"].(string)
									}

									if vvv,ok := m2["type"]; vvv == "drsg" &&ok{
										fmt.Println(m2["txt"])
										returnData.DressIndex.Txt = m2["txt"].(string)
										returnData.DressIndex.Brf = m2["brf"].(string)
									}

									if vvv,ok := m2["type"]; vvv == "flu" &&ok{
										fmt.Println(m2["txt"])
										returnData.ColdIndex.Txt = m2["txt"].(string)
										returnData.ColdIndex.Brf = m2["brf"].(string)
									}

									if vvv,ok := m2["type"]; vvv == "sport" &&ok{
										fmt.Println(m2["txt"])
										returnData.Sport.Txt = m2["txt"].(string)
										returnData.Sport.Brf = m2["brf"].(string)
									}

									if vvv,ok := m2["type"]; vvv == "trav" &&ok{
										fmt.Println(m2["txt"])
										returnData.Travel.Txt = m2["txt"].(string)
										returnData.Travel.Brf = m2["brf"].(string)
									}

									if vvv,ok := m2["type"]; vvv == "uv" &&ok{
										fmt.Println(m2["txt"])
										returnData.Uv.Txt = m2["txt"].(string)
										returnData.Uv.Brf = m2["brf"].(string)
									}

									if vvv,ok := m2["type"]; vvv == "cw" &&ok{
										fmt.Println(m2["txt"])
										returnData.WashCarIndex.Txt = m2["txt"].(string)
										returnData.WashCarIndex.Brf = m2["brf"].(string)
									}

									if vvv,ok := m2["type"]; vvv == "air" &&ok{
										fmt.Println(m2["txt"])
										returnData.AirPollution.Txt = m2["txt"].(string)
										returnData.AirPollution.Brf = m2["brf"].(string)
									}

									//for k,v := range m2{
									//	fmt.Printf("result %T %T\n",k,v)
									//	if k == "type" && v == "comf"{
									//		fmt.Println("comf")
									//
									//	}
									//}

								}
							}
						}

					}
					//m2 := make(map[string]interface{})
					//m2 = m1[3].(map[string]interface{})
				}
			}

		}
		//m1 = data

	}


	p.WeatherData = returnData
	//return returnData
}
