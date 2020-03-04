package hefeng

import (
	"encoding/json"
	"fmt"
)

func (p *HeFWeather) Parse(contents []byte) bool {
	var returnData WeatherInfo

	m := make(map[string]interface{})
	json.Unmarshal(contents, &m)

	if value, ok := m["HeWeather6"]; ok {
		if len(value.([]interface{})) == 0 {
			return false
		}
		u := value.([]interface{})[0]
		if value1, ok := u.(map[string]interface{})["lifestyle"]; ok {
			if len(value1.([]interface{})) == 0 {
				return false
			}
			for i := range value1.([]interface{}) {
				value2 := make(map[string]interface{})
				value2 = value1.([]interface{})[i].(map[string]interface{})
				if vvv, ok := value2["type"]; vvv == "comf" && ok {
					fmt.Println(value2["txt"])
					returnData.Comfort.Txt = value2["txt"].(string)
					returnData.Comfort.Brf = value2["brf"].(string)
				}

				if vvv, ok := value2["type"]; vvv == "drsg" && ok {
					fmt.Println(value2["txt"])
					returnData.DressIndex.Txt = value2["txt"].(string)
					returnData.DressIndex.Brf = value2["brf"].(string)
				}

				if vvv, ok := value2["type"]; vvv == "flu" && ok {
					fmt.Println(value2["txt"])
					returnData.ColdIndex.Txt = value2["txt"].(string)
					returnData.ColdIndex.Brf = value2["brf"].(string)
				}

				if vvv, ok := value2["type"]; vvv == "sport" && ok {
					fmt.Println(value2["txt"])
					returnData.Sport.Txt = value2["txt"].(string)
					returnData.Sport.Brf = value2["brf"].(string)
				}

				if vvv, ok := value2["type"]; vvv == "trav" && ok {
					fmt.Println(value2["txt"])
					returnData.Travel.Txt = value2["txt"].(string)
					returnData.Travel.Brf = value2["brf"].(string)
				}

				if vvv, ok := value2["type"]; vvv == "uv" && ok {
					fmt.Println(value2["txt"])
					returnData.Uv.Txt = value2["txt"].(string)
					returnData.Uv.Brf = value2["brf"].(string)
				}

				if vvv, ok := value2["type"]; vvv == "cw" && ok {
					fmt.Println(value2["txt"])
					returnData.WashCarIndex.Txt = value2["txt"].(string)
					returnData.WashCarIndex.Brf = value2["brf"].(string)
				}

				if vvv, ok := value2["type"]; vvv == "air" && ok {
					fmt.Println(value2["txt"])
					returnData.AirPollution.Txt = value2["txt"].(string)
					returnData.AirPollution.Brf = value2["brf"].(string)
				}
			}
		} else {
			return false
		}
	} else {
		return false
	}

	p.WeatherData = returnData
	return true
}
