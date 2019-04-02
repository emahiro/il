package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		panic(err)
	}
}
