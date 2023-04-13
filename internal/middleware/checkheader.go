package middleware

import (
	"log"
	"net/http"
)

func CheckHeader(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("CheckHeader")
		if r.Header.Get("Content-Type") != "application/json" {
			rw.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		next(rw, r)
	}
}
