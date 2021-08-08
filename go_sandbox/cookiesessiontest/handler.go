package main

import (
	"fmt"
	"net/http"
)

var (
	ckName   = "test"
	keyPairs = []byte("test")
)

type Handler struct{}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("OK.")))
}
