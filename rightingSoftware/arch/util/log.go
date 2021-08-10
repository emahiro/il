package util

import (
	"log"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v", r.Method, r.URL)
		next.ServeHTTP(rw, r)
	})
}
