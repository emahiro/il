package main

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func main() {
	mux := http.NewServeMux()
	h := &Handler{}
	mux.HandleFunc("get", h.Get)
	if err := http.ListenAndServe(":8080", SessionCookie(&sessions.Options{}, Logger(mux))); err != nil {
		panic(err)
	}
}
