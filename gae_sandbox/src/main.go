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

	"emahiro/il/gae_sandbox/router"
)

var addr = 8080

func main() {
	r := router.NewRouter()
	handler := r.Build(router.PureServeMux())

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", addr),
		Handler: handler,
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
