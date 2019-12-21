package handler

import (
	"html/template"
	"net/http"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("hello appengine go1.11 world!!!")); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// RenderHandler is render tmpl file
func RenderHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(w, nil); err != nil {
		panic(err)
	}
}
