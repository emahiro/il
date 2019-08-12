package router

import (
	"net/http"

	"github.com/emahiro/ae-plain-logger/middleware"
	"github.com/fharding1/gemux"
	"github.com/go-chi/chi"

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

	return middleware.MwAEPlainLogger("NewServeMux")(mux)
}

func (r *WebRouter) Gemux() http.Handler {
	mux := &gemux.ServeMux{}
	mux.Handle("/hello", http.MethodGet, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	}))
	mux.Handle("/metadata", http.MethodGet, http.HandlerFunc(metadata.GetMetadata))
	mux.Handle("/verify", http.MethodGet, http.HandlerFunc(metadata.Verify))

	return middleware.MwAEPlainLogger("GemuxServeMux")(mux)
}

func (r *WebRouter) ChiMux() http.Handler {
	mux := chi.NewMux()
	mux.HandleFunc("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	}))
	mux.Handle("/metadata", http.HandlerFunc(metadata.GetMetadata))
	mux.Handle("/verify", http.HandlerFunc(metadata.Verify))

	return middleware.MwAEPlainLogger("chiServeMux")(mux)
}
