package handler

import (
	"net/http"
)

// Index ...
func Index(w http.ResponseWriter,r *http.Request){

	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("hello appengine go1.11 world!!!")); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
