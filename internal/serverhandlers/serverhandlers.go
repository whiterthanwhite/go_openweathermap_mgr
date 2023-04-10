package serverhandlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func AddWeatherMeasurement() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		fmt.Println(string(reqBody))
		rw.WriteHeader(http.StatusOK)
	}
}
