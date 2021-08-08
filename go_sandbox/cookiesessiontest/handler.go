package main

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	ck, err := r.Cookie("test")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("OK. cookie value is %#v", ck)))
}
