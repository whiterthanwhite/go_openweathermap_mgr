package currentweather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type CurrentWeather interface {
	GetByGeoCoordinates(lat, lon float32) error
}

type CurrentWeatherData struct {
	APIKey string
}

func (cwd *CurrentWeatherData) GetByGeoCoordinates(lat, lon float32) (*CurrentWeatherResponse, error) {
	return cwd.doRequest(lat, lon)
}

// url := `https://api.openweathermap.org/data/2.5/weather?lat={lat}&lon={lon}&appid={API key}`
// Possible values
// lat, lon (required)
// appid (required)
// mode (optional) (json default)
// units (optional)
// lang (optional)
func (cwd *CurrentWeatherData) doRequest(lat, lon float32) (*CurrentWeatherResponse, error) {
	urlString, err := cwd.getURL(lat, lon)
	if err != nil {
		return nil, err
	}

	r, err := http.Get(urlString)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	currWeatherResp, err := parseResponse(r.Body)
	if err != nil {
		return nil, err
	}

	return currWeatherResp, nil
}

func (cwd *CurrentWeatherData) getURL(lat, lon float32) (string, error) {
	u, err := url.Parse(`https://api.openweathermap.org/data/2.5/weather`)
	if err != nil {
		return "", err
	}
	if cwd.APIKey == "" {
		return "", errors.New("api key is not implemented")
	}
	urlValues := url.Values{}
	urlValues.Add("lat", fmt.Sprint(lat))
	urlValues.Add("lon", fmt.Sprint(lon))
	urlValues.Add("units", "metric")
	urlValues.Add("appid", cwd.APIKey)
	return fmt.Sprintf("%s?%s", u, urlValues.Encode()), nil
}

func parseResponse(r io.Reader) (*CurrentWeatherResponse, error) {
	c := &CurrentWeatherResponse{}
	jsonDecoder := json.NewDecoder(r)
	if err := jsonDecoder.Decode(c); err != nil {
		return nil, err
	}
	return c, nil
}
