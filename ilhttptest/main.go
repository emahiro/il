package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const port = ":8080"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	})

	server := http.Server{
		Addr:    port,
		Handler: mux,
	}

	fmt.Printf("start server port: %v ....\n", port)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("server closed with error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	fmt.Printf("signal %d received, server shutdown ...\n", <-quit)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("failed to graceful shutdown. err: %v", err)
		os.Exit(1)
	}

	fmt.Println("server shutdown....")
	os.Exit(0)
}
