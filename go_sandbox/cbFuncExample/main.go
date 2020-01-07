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
	"strings"
	"syscall"
	"time"
)

var port = 8080

type user struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	})
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		log.Println("get user")
		u := user{ID: 1, Name: "Taro"}

		buf := bytes.NewBuffer([]byte(""))
		enc := json.NewEncoder(buf)
		enc.SetIndent("", strings.Repeat(" ", 2))
		if err := enc.Encode(&u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("parse error. err: %v", err)))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(buf.Bytes())
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	log.Println("start server....")
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalln("cannot start server.")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	log.Printf("signal %d received, server shutdown ...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("faild to graceful shutdown. err: %v", err)
	}

	log.Println("success to shutdown server")
	os.Exit(0)
}
