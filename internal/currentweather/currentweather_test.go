package currentweather

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestGetByGeoCoordinates(t *testing.T) {
	t.Fail()
}

func TestGetURL(t *testing.T) {
	c := &CurrentWeatherData{}
	urlStr, err := c.getURL(50.0, 20.0)
	if err == nil {
		t.Error("Should be error because of key absence")
	}
	c = &CurrentWeatherData{
		APIKey: "123456789",
	}
	urlStr, err = c.getURL(50.4005, 20.5004)
	if err != nil {
		t.Error(err)
	}
	expectedURL := `https://api.openweathermap.org/data/2.5/weather?appid=123456789&lat=50.4005&lon=20.5004`
	if !strings.Contains(urlStr, expectedURL) {
		t.Errorf("wrong url\nactual: %s\nexpected: %s\n", urlStr, expectedURL)
	}
}

func TestParseResponse(t *testing.T) {
	input := getMockData()
	_, err := parseResponse(input)
	if err != nil {
		t.Error(err.Error())
	}
}

func getMockData() io.Reader {
	return bytes.NewReader([]byte(`{"coord":{"lon":13.4,"lat":52.52},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"base":"stations","main":{"temp":280.03,"feels_like":278.22,"temp_min":278.2,"temp_max":280.94,"pressure":1010,"humidity":55},"visibility":10000,"wind":{"speed":2.57,"deg":130},"clouds":{"all":0},"dt":1680716355,"sys":{"type":2,"id":2009543,"country":"DE","sunrise":1680669137,"sunset":1680716747},"timezone":7200,"id":6545310,"name":"Mitte","cod":200}`))
}
