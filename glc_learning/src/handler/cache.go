package handler

import (
	"net/http"
)

// SetCache ...
func SetCache(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("set cache!"))
}

// GetCache ...
func GetCache(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("get cache!"))
}
