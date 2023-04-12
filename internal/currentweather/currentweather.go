package currentweather

import (
	"bytes"
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

func (cwd *CurrentWeatherData) SendDataToServer(cwr *CurrentWeatherResponse) error {
	urlString, err := url.Parse("http://localhost:8080/currentweather")
	if err != nil {
		return nil
	}

	reqBody, err := json.Marshal(cwr)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodGet, urlString.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnsupportedMediaType {
		return errors.New("unsupported media type")
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("actual status: %d; expected status: %d; store failed!",
			resp.StatusCode, http.StatusOK)
	}

	fmt.Println("store succeed!")
	return nil
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
