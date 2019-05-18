package http

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func LoggingHandlerFunc(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lw := loggingWriter{w, 200}
		if next != nil {
			next.ServeHTTP(lw, r)
		}
		log.
			WithField("type", "access").
			WithField("path", r.URL.Path).
			WithField("host", r.Host).
			WithField("remote", r.RemoteAddr).
			WithField("response_code", lw.statusCode).
			Infoln("Received a request")
	}
}

type loggingWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w loggingWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
