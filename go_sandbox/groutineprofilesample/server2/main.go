package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var port = 8081

type user struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
		return
	})
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		user := user{Name: "Taro", Age: 12}

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(&user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(buf.Bytes())
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	log.Println("server start....")
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	log.Printf("signal %d received. server begins to shutdown...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("failed to shutdown gracefully")
	}

	log.Println("success to shutdown")
	os.Exit(0)
}
