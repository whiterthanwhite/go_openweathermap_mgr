package serverhandlers

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/whiterthanwhite/go_openweathermap_mgr/internal/currentweather"
	"github.com/whiterthanwhite/go_openweathermap_mgr/internal/database"
)

func AddWeatherMeasurement() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("AddWeatherMeasurement")
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		weather, err := parseBody(reqBody)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), time.Second)
		defer cancel()

		db := database.GetInstance()
		if err = db.CreateConnection(ctx); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err = db.InsertWeatherData(ctx, weather); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
	}
}

func parseBody(reqBody []byte) (*currentweather.Weather, error) {
	if len(reqBody) == 0 {
		return nil, errors.New("body is empty")
	}
	weatherResponse := &currentweather.CurrentWeatherResponse{}
	if err := json.Unmarshal(reqBody, weatherResponse); err != nil {
		return nil, err
	}
	w := &currentweather.Weather{
		Lon:  weatherResponse.Coord.Lon,
		Lat:  weatherResponse.Coord.Lat,
		Temp: weatherResponse.Main.Temp,
	}
	return w, nil
}
