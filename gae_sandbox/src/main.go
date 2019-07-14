package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fharding1/gemux"

	"emahiro/il/gae_sandbox/model"
)

var addr = 8080

func main() {
	mux := &gemux.ServeMux{}
	mux.Handle("/hello", http.MethodGet, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	}))
	mux.Handle("/metadata", http.MethodGet, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := http.Client{Transport: http.DefaultTransport}

		values := url.Values{}
		values.Set("audience", "test")
		values.Set("format", "full")

		reqURL := fmt.Sprintf("http://metadata/computeMetadata/v1/instance/service-accounts/default/identity?%s", values.Encode())
		req, err := http.NewRequest(http.MethodGet, reqURL, nil)
		if err != nil {
			log.Printf("failed to create request")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		req.Header.Set("Metadata-Flavor", "Google")

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("failed to get metadata. err: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		w.WriteHeader(resp.StatusCode)

		if resp.StatusCode != http.StatusOK {
			log.Printf("failed to get metadata. body: %+v", resp.Body)
			return
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("failed to read body. err: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Write(b)
	}))
	mux.Handle("/verify", http.MethodGet, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := http.Client{Transport: http.DefaultTransport}
		req, err := http.NewRequest(http.MethodGet, "https://www.googleapis.com/oauth2/v3/certs", nil)
		if err != nil {
			log.Printf("failed to create request. err: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("failed to get public keys. err: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		w.WriteHeader(resp.StatusCode)

		if resp.StatusCode != http.StatusOK {
			log.Printf("failed to get public kyes. code: %d", resp.StatusCode)
			return
		}

		keys := model.PublicKeys{}
		if err := json.NewDecoder(resp.Body).Decode(&keys); err != nil {
			log.Printf("failed to decode json. err: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// verify

	}))

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
