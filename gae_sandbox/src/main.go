package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fharding1/gemux"

	"emahiro/il/gae_sandbox/metadata"
)

var addr = 8080

func main() {
	mux := &gemux.ServeMux{}
	mux.Handle("/hello", http.MethodGet, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	}))
	mux.Handle("/metadata", http.MethodGet, http.HandlerFunc(metadata.GetMetadata))
	mux.Handle("/verify", http.MethodGet, http.HandlerFunc(metadata.Verify))

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", addr),
		Handler: mux,
	}

	log.Printf("server start port: %d ...", addr)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("server closed with error: %v", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	log.Printf("SIGNAL %d received, server shutdown ...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("failed to gracefully shutdown. err: %v", err)
		os.Exit(1)
	}
	log.Printf("server shutdown....")
}
