package router

import (
	"net/http"

	"github.com/emahiro/il/glc_learning/src/handler"
)

// Router is ...
type Router interface {
	Build(h http.Handler) http.Handler
}

// WebRouter is ...
type WebRouter struct{}

func (r *WebRouter) ServeHTTP(res http.ResponseWriter, req *http.Request) {}

// Build is ...
func (r *WebRouter) Build(h http.Handler) http.Handler {
	return h
}

// DefaultRouter is ...
func DefaultRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	})
	mux.HandleFunc("/get", handler.GetCache)
	mux.HandleFunc("/set", handler.SetCache)

	return mux
}
