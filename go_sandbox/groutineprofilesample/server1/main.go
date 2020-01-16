package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var port = 8080

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
		return
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

type user struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func getUser() {
	client := http.DefaultClient
	req, _ := http.NewRequest("GET", "http://localhost:8081/users", nil)

	cnt := 25
	wg := sync.WaitGroup{}
	wg.Add(cnt)

	for i := 0; i < cnt; i++ {
		go func() {
			defer wg.Done()
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			var user user
			if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
				panic(err)
			}
		}()
	}

	wg.Wait()
}
