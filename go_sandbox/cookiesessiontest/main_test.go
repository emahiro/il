package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	ret := m.Run()
	os.Exit(ret)
}

func TestRouter(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		h := &Handler{}
		h.Get(rw, r)
	}))
	t.Cleanup(ts.Close)

	t.Log(ts.URL)

	req, _ := http.NewRequest(http.MethodGet, ts.URL, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", string(b))
}
