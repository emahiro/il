package handler

import (
	"net/http"

	"github.com/emahiro/ae-plain-logger/log"
	"github.com/emahiro/glc"
)

var fc *glc.FileCache

func init() {
	var err error
	if fc, err = glc.NewFileCache("glc_learning"); err != nil {
		panic(err)
	}
}

// SetCache ...
func SetCache(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := fc.Set("test", []byte("this is cache")); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Infof(ctx, "failed to set error. err: %v", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("set cache!"))
}

// GetCache ...
func GetCache(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("get cache!"))
}
