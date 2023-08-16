package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

var addr = 8080

func main() {
	ctx := context.Background()
	run(ctx)
}

func run(ctx context.Context) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	})

	server := &http.Server{
		Addr:    ":" + fmt.Sprint(addr),
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		if err := server.Shutdown(ctx); err != nil {
			panic(err)
		}
	}()

	log.Println("start server at port: " + fmt.Sprint(addr))
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
