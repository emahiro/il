package router

import (
	"net/http"

	"github.com/fharding1/gemux"

	"emahiro/il/gae_sandbox/metadata"
)

type Router interface {
	Build() http.Handler
}

type WebRouter struct{}

func NewRouter() *WebRouter {
	return &WebRouter{}
}

func (r *WebRouter) NewServeMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	}))
	mux.Handle("/metadata", http.HandlerFunc(metadata.GetMetadata))
	mux.Handle("/verify", http.HandlerFunc(metadata.Verify))

	return mux
}

func (r *WebRouter) Gemux() http.Handler {
	mux := &gemux.ServeMux{}
	mux.Handle("/hello", http.MethodGet, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	}))
	mux.Handle("/metadata", http.MethodGet, http.HandlerFunc(metadata.GetMetadata))
	mux.Handle("/verify", http.MethodGet, http.HandlerFunc(metadata.Verify))

	return mux
}
