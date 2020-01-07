package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

var port = 8080

type user struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type cache struct {
	data map[string][]byte
	exp  int64
	mu   sync.Mutex
}

var c = &cache{data: make(map[string][]byte)}

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
	mux.HandleFunc("/cache", func(w http.ResponseWriter, r *http.Request) {
		f := func() ([]byte, error) {
			return getUser()
		}
		b, err := c.getAndSetAfterFuncCall("test", f)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error: %v", err)))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
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

func getUser() ([]byte, error) {
	cli := http.DefaultClient
	resp, err := cli.Get("http://localhost:8080/user")
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("reqest failed. code: %d", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (c *cache) getAndSetAfterFuncCall(key string, f func() ([]byte, error)) ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c == nil {
		return nil, fmt.Errorf("error: attempting to access nil cache object")
	}

	if len(key) == 0 {
		return nil, fmt.Errorf("error: key is empty")
	}

	if c.data == nil {
		return nil, fmt.Errorf("error: attempting to access nil map")
	}

	if time.Now().Unix() < c.exp {
		if data, ok := c.data[key]; ok {
			log.Println("use local cache")
			return data, nil
		}
	}

	b, err := f()
	if err != nil {
		return nil, err
	}

	c.data[key] = b
	c.exp = time.Now().Add(10 * time.Second).Unix()

	return b, nil
}
