package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
