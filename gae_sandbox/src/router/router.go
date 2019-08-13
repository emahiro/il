package router

import (
	"net/http"

	"github.com/emahiro/ae-plain-logger/middleware"
	"github.com/fharding1/gemux"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi"

	"emahiro/il/gae_sandbox/metadata"
)

type Router interface {
	Build(mux http.Handler) http.Handler
}

type WebRouter struct{}

func (r *WebRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {}

func NewRouter() Router {
	return &WebRouter{}
}

func (r *WebRouter) Build(mux http.Handler) http.Handler {
	return middleware.MwAEPlainLogger("MwAEPlainLogger")(r)
}

func PureServeMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	}))

	return mux
}

func Gemux() http.Handler {
	mux := &gemux.ServeMux{}
	mux.Handle("/hello", http.MethodGet, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	}))
	mux.Handle("/metadata", http.MethodGet, http.HandlerFunc(metadata.GetMetadata))
	mux.Handle("/verify", http.MethodGet, http.HandlerFunc(metadata.Verify))

	return mux
}

func ChiMux() http.Handler {
	mux := chi.NewMux()
	mux.HandleFunc("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	}))
	mux.Handle("/metadata", http.HandlerFunc(metadata.GetMetadata))
	mux.Handle("/verify", http.HandlerFunc(metadata.Verify))

	return mux
}

func GinRouter() http.Handler {
	mux := gin.New()
	mux.GET("/hello", func(gc *gin.Context) {
		gc.Status(http.StatusOK)
		gc.Writer.Write([]byte("hello"))
		return
	})

	return mux
}
