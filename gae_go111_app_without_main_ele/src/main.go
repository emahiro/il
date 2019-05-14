package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// main.goの中身
func main() {
	f, err := os.Open("app.yaml")
	log.Printf("%v", err)
	if _, err := io.Copy(os.Stdout, f); err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		panic(err)
	}
}
