package main

import "net/http"

import "fmt"

func mw() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// middleware で何かしらの処理をする
			fmt.Printf("This is in middleware\n")
		})
	}
}
