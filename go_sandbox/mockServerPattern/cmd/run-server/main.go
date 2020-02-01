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
)

var port = 8081

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("enexpected request"))
			return
		}
		fmt.Println("test")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("hello at port: %d", port)))
	})

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	m := <-quit
	log.Printf("received signal %v. start shutdown server...", m)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown gracefully. err: %v", err)
	}

	log.Println("success to shutdown gracefully.")
	os.Exit(0)
}
