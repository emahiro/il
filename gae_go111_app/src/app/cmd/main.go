package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"gae_go111_app/handler"
)

func main() {
	f, err := os.Open("app.yaml")
	log.Printf("%v", err)
	io.Copy(os.Stdout, f)

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
