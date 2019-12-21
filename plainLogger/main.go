package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/emahiro/log_output/mw"
)

var port = 8080

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mw.Logger()(mux),
	}

	fmt.Printf("start server in port: %v...\n", port)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Printf("failed to start server. err: %v\n", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	fmt.Printf("shutdown server... signal: %v\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("failed to graceful shutdown. err: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("server shutdown.")
	os.Exit(0)
}
