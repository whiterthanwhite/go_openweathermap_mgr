package middleware

import (
	"fmt"
	"io"
	"net/http"
)

type Logger struct {
	Handler   http.Handler
	LogWriter io.Writer
}

func (l *Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(l.LogWriter, "Request host: %s\n", r.Host)
	fmt.Fprintf(l.LogWriter, "Request uri: %s\n", r.RequestURI)
	fmt.Fprintf(l.LogWriter, "Request method: %s\n", r.Method)

	l.Handler.ServeHTTP(rw, r)
}
