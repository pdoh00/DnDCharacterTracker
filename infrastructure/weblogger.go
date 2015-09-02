package infrastructure

import (
	"log"
	"net/http"
	"time"
)

//WebLogger http handler decorator which logs method, uri, name, duration
func WebLogger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)
		log.Printf("%s\t%s\t%s\t%s\t",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start))
	})
}
