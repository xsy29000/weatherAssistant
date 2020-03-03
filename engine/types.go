package engine

//type Parser interface {
//	Parse([]byte) WeatherInfo
//}

type Inform interface {
	Inform(string,string) bool
}

type Request struct {
	Url    string
}

type Display interface {
	DisplayData(interface{}) string
}