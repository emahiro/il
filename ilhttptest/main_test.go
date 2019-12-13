package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var ts *httptest.Server

func spinUp() *httptest.Server {
	return httptest.NewServer(mw()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// この中で何かしらテストをしたいとき
		fmt.Printf("request in test server. req: %+v\n", r)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello in test server."))
	})))
}

func TestMain(m *testing.M) {
	ts = spinUp()
	defer ts.Close()

	ret := m.Run()
	os.Exit(ret)
}

func TestRequest(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ts.Config.Handler.ServeHTTP(w, r)
}
