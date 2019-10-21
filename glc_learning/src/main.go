package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/emahiro/il/glc_learning/router"
)

func main() {
	r := &router.WebRouter{}
	h := r.Build(router.DefaultRouter())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("start server in :%s port ...", port)

	server := http.Server{
		Addr:    ":" + port,
		Handler: h,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("failed to start server. err: %v\n", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	log.Printf("shutdown server... signal: %v\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("failed to shutdown gracefully ...\n")
		os.Exit(1)
	}

	log.Printf("server shutdown ...\n")
	os.Exit(0)
}
