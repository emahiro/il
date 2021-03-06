package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var port = 8080

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &handler{})

	server := &http.Server{
		Handler: mux,
		Addr:    fmt.Sprintf(":%d", port),
	}

	log.Println("start server....")
	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	log.Printf("syscall %d received. start to shutdown server...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown gracefully, err: %v", err)
	}

	log.Println("success to shutdown server.")
	os.Exit(0)
}

type handler struct{}

type req struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(d))

	w.WriteHeader(http.StatusOK)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	req := &req{}
	if err := json.Unmarshal(b, req); err != nil {
		log.Fatal(err)
	}
	fmt.Println(req)
	w.Write([]byte("ok"))
}
