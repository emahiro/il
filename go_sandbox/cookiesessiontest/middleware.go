package main

import (
	"fmt"
	"log"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		line := fmt.Sprintf("%s %s", r.Method, r.URL.Path)
		log.Println(line)
		next.ServeHTTP(w, r)
	})
}
