package handler

import (
	"emahiro/gotester/model"
	"encoding/json"
	"log"
	"net/http"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("index"))
}

// GetUserA ...
func GetUserA(w http.ResponseWriter, r *http.Request) {
	u, err := model.GetUserA()
	if err != nil {
		log.Fatalf("failed to get UserA resources. err: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(u)
	if err != nil {
		log.Fatalf("failed to encode json. err: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

// GetUserB ...
func GetUserB(w http.ResponseWriter, r *http.Request) {}
