package fetcher

import (
	"fmt"
	"testing"
)

func TestSave(t *testing.T) {

	var url string = "https://free-api.heweather.net/s6/weather/lifestyle?location=guangzhou&key=26e6e05426344f8da738eb82bd30e574"

	data, err := Fetch(url)

	fmt.Println(string(data))

	if err != nil {
		t.Errorf("got %v", err)
	}

}
