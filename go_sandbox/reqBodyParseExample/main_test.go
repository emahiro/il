package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	g, err := ioutil.ReadFile(filepath.Join("testdata", t.Name()+".golden"))
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d, _ := httputil.DumpRequest(r, true)
		t.Log(string(d))

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		p, err := strconv.Unquote(string(b))
		if err != nil {
			t.Fatal(err)
		}
		var req req
		if err := json.Unmarshal([]byte(p), &req); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	s := strings.Trim(string(g), "\n")
	req, _ := http.NewRequest(http.MethodPost, ts.URL, bytes.NewBuffer([]byte(s)))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
}
