package server

import (
	"log"
	"net/http"
	"time"
)

func CreateLoggerAdapter(name string) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			h.ServeHTTP(w, r)
			log.Printf(
				"%s\t%s",
				name,
				time.Since(start),
			)
		})
	}
}
