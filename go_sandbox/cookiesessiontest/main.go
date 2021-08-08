package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	h := &Handler{}
	mux.HandleFunc("get", h.Get)
	if err := http.ListenAndServe(":8080", Logger(mux)); err != nil {
		panic(err)
	}
}
