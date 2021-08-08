package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
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

	store := sessions.NewCookieStore(keyPairs)
	store.Options = &sessions.Options{}
	session, err := store.Get(r, ckName)
	if err != nil {
		// refresh session
		session = sessions.NewSession(store, ckName)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("OK. cookie name is %s value is %#v", session.Name(), session.Values["test"])))
}
