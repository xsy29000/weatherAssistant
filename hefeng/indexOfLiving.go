package hefeng

import (
	"encoding/json"
	"fmt"
    "strconv"
)

func (p *HeFWeather) Parse(contents []byte) bool {
	var returnData WeatherInfo

	m := make(map[string]interface{})
	json.Unmarshal(contents, &m)

	if value, ok := m["daily"]; ok {                                                                                                        
 		for i := range value.([]interface{}){
			value2 := make(map[string]interface{})
			value2 = value.([]interface{})[i].(map[string]interface{})
                   
            infoType,err := strconv.Atoi(value2["type"].(string))
            if err != nil{
                return false
            }
			switch infoType{
                case 8:{
					fmt.Println(value2["text"])
					returnData.Comfort.Txt = value2["text"].(string)
					returnData.Comfort.Brf = value2["category"].(string)
				}

                case 3:{
					fmt.Println(value2["text"])
					returnData.DressIndex.Txt = value2["text"].(string)
					returnData.DressIndex.Brf = value2["category"].(string)
				}

                case 9:{
					fmt.Println(value2["text"])
					returnData.ColdIndex.Txt = value2["text"].(string)
					returnData.ColdIndex.Brf = value2["category"].(string)
				}

                case 1:{
					fmt.Println(value2["text"])
					returnData.Sport.Txt = value2["text"].(string)
					returnData.Sport.Brf = value2["category"].(string)
				}

                case 6:{
					fmt.Println(value2["text"])
					returnData.Travel.Txt = value2["text"].(string)
					returnData.Travel.Brf = value2["category"].(string)
				}

                case 5:{
					fmt.Println(value2["text"])
					returnData.Uv.Txt = value2["text"].(string)
					returnData.Uv.Brf = value2["category"].(string)
				}

                case 2:{
					fmt.Println(value2["text"])
					returnData.WashCarIndex.Txt = value2["text"].(string)
					returnData.WashCarIndex.Brf = value2["category"].(string)
				}
	        
                case 10:{
					fmt.Println(value2["text"])
					returnData.AirPollution.Txt = value2["text"].(string)
					returnData.AirPollution.Brf = value2["category"].(string)
				}
            }
	    }
    }

/*
	if value, ok := m["daily"]; ok {
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
					fmt.Println(value2["text"])
					returnData.Comfort.Txt = value2["text"].(string)
					returnData.Comfort.Brf = value2["category"].(string)
				}

				if vvv, ok := value2["type"]; vvv == "drsg" && ok {
					fmt.Println(value2["text"])
					returnData.DressIndex.Txt = value2["text"].(string)
					returnData.DressIndex.Brf = value2["category"].(string)
				}

				if vvv, ok := value2["type"]; vvv == "flu" && ok {
					fmt.Println(value2["text"])
					returnData.ColdIndex.Txt = value2["text"].(string)
					returnData.ColdIndex.Brf = value2["category"].(string)
				}

				if vvv, ok := value2["type"]; vvv == "sport" && ok {
					fmt.Println(value2["text"])
					returnData.Sport.Txt = value2["text"].(string)
					returnData.Sport.Brf = value2["category"].(string)
				}

				if vvv, ok := value2["type"]; vvv == "trav" && ok {
					fmt.Println(value2["text"])
					returnData.Travel.Txt = value2["text"].(string)
					returnData.Travel.Brf = value2["category"].(string)
				}

				if vvv, ok := value2["type"]; vvv == "uv" && ok {
					fmt.Println(value2["text"])
					returnData.Uv.Txt = value2["text"].(string)
					returnData.Uv.Brf = value2["category"].(string)
				}

				if vvv, ok := value2["type"]; vvv == "cw" && ok {
					fmt.Println(value2["text"])
					returnData.WashCarIndex.Txt = value2["text"].(string)
					returnData.WashCarIndex.Brf = value2["category"].(string)
				}

				if vvv, ok := value2["type"]; vvv == "air" && ok {
					fmt.Println(value2["text"])
					returnData.AirPollution.Txt = value2["text"].(string)
					returnData.AirPollution.Brf = value2["category"].(string)
				}
			}
		} else {
			return false
		}
	} else {
		return false
	}
*/
	p.WeatherData = returnData
	return true
}
