package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(m *testing.M) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// この中で何かしらテストをしたいとき
		fmt.Printf("request in test server. req: %+v", r)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello in test server."))
	}))
	defer ts.Close()

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ts.Config.Handler.ServeHTTP(w, r)
}
