package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestPlainServerMock(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test"))
	}))
	defer ts.Close()

	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res != nil && res.StatusCode != http.StatusOK {
		t.Fatalf("err: status code: %v", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", string(b))
}

func Test_userHandler(t *testing.T) {
	ts := httptest.NewServer(&userHandler{})
	defer ts.Close()

	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("err: status code is %d", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", string(b))
}

func Test_userHandlerMock(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("user handler test"))
	}))
	defer ts.Close()

	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("err: status code is %d", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", string(b))
}

func TestAnotherServerMock(t *testing.T) {
	httpmock.RegisterResponder(http.MethodGet, "/", func(r *http.Request) (*http.Response, error) {
		r.URL.Scheme = "http"
		r.URL.Host = "localhost:8081"
		r.URL.Path = ""
		return http.DefaultTransport.RoundTrip(r)
	})

	client := &http.Client{
		Transport: httpmock.DefaultTransport,
	}
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("err: status code is %d", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", string(b))
}

func TestAnotherServerMockWithTestServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("unexpected method"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test server response is ok"))
	}))
	defer ts.Close()

	httpmock.RegisterResponder(http.MethodGet, "/", func(r *http.Request) (*http.Response, error) {
		u, err := url.Parse(ts.URL)
		if err != nil {
			t.Fatal(err)
		}
		r.URL = u
		return http.DefaultTransport.RoundTrip(r)
	})

	client := &http.Client{
		Transport: httpmock.DefaultTransport,
	}
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("err: status code is %d", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", string(b))
}

func TestAnotherServerMockWithTestServerURLByHTTPDefaultTransport(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("unexpected method"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test server response is ok via http default transport in httpmock"))
	}))
	defer ts.Close()

	httpmock.RegisterResponder(http.MethodGet, "/", func(r *http.Request) (*http.Response, error) {
		client := &http.Client{
			Transport: new(http.Transport),
		}
		url, err := url.Parse(ts.URL)
		if err != nil {
			return nil, err
		}
		r.URL = url
		return client.Do(r)
	})

	client := &http.Client{
		Transport: httpmock.DefaultTransport,
	}
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("err: status code is %d", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", string(b))
}

func TestAnotherServerMockWithTestServerTrasport(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("unexpected method"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test server response is ok in using test server http default transport"))
	}))
	defer ts.Close()

	httpmock.RegisterResponder(http.MethodGet, ts.URL, func(r *http.Request) (*http.Response, error) {
		clinet := &http.Client{
			Transport: new(http.Transport),
		}
		return clinet.Do(r)
	})

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("err: status code is %d", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", string(b))
}
