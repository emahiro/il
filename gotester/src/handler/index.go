package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"emahiro/gotester/model"
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
		return
	}

	b, err := json.Marshal(u)
	if err != nil {
		log.Fatalf("failed to encode json. err: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// GetUserB ...
func GetUserB(w http.ResponseWriter, r *http.Request) {
	u := &model.UserB{Name: "Alice", Age: 1}
	uu, err := u.Get()
	if err != nil {
		log.Fatalf("failed to get user b resouces")
		return
	}

	log.Printf("uu is %v", uu)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("uu is %#v", uu)))
}
