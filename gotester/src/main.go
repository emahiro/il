package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fharding1/gemux"

	"emahiro/gotester/handler"
)

const (
	port = ":8080"
)

func main() {
	mux := &gemux.ServeMux{}
	mux.Handle("/", http.MethodGet, http.HandlerFunc(handler.Index))
	mux.Handle("/userA", http.MethodGet, http.HandlerFunc(handler.GetUserA))
	mux.Handle("/userB", http.MethodGet, http.HandlerFunc(handler.GetUserB))

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("server closed with %v", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	log.Printf("SIGNAL %v received, then shutting down...\n", <-quit)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to graceful shutdown: %v", err)
	}
	log.Printf("server shutdown.")
}
