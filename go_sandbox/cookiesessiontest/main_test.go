package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
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

	req, _ := http.NewRequest(http.MethodGet, ts.URL, nil)
	codec := securecookie.CodecsFromPairs(keyPairs)
	ckstr, err := codec[0].Encode("test", map[interface{}]interface{}{
		"test": "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(&http.Cookie{
		Name:  "test",
		Value: ckstr,
	})

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

func TestRouterWithMiddleware(t *testing.T) {
	codec := securecookie.CodecsFromPairs(keyPairs)
	ckstr, err := codec[0].Encode(ckName, map[interface{}]interface{}{
		"test": "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodGet, "http://test.com/get", nil)
	req.AddCookie(&http.Cookie{
		Name:  ckName,
		Value: ckstr,
	})

	rec := httptest.NewRecorder()
	mux := http.NewServeMux()
	h := &Handler{}
	mux.HandleFunc("/get", h.Get)
	SessionCookie(&sessions.Options{}, mux).ServeHTTP(rec, req)
}
