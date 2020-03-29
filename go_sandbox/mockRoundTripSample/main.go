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

	"github.com/emahiro/hbhcl"

	"github.com/emahiro/il/go_sandbox/mockRoundTripSample/mw"
)

var port = 8080

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(http.StatusText(http.StatusOK)))
	})
	mux.HandleFunc("/hatena", func(w http.ResponseWriter, r *http.Request) {
		feed, err := hbhcl.NewClient("ema-hiro").FetchFeed()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(feed.Entries[0].Title))
	})

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mw.Logger(mux),
	}

	log.Printf("starting server at %d ...\n", port)

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatalf("failed to start server. err: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("%+v received, start to shutdown server gracefully\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown server gracefully. err: %v", err)
	}

	log.Println("success to shutdown server.")
	os.Exit((0))
}
