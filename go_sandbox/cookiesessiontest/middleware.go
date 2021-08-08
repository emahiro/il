package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		line := fmt.Sprintf("%s %s", r.Method, r.URL.Path)
		log.Println(line)
		next.ServeHTTP(w, r)
	})
}

func SessionCookie(opts *sessions.Options, next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		store := sessions.NewCookieStore(keyPairs)
		store.Options = opts
		session, err := store.Get(r, ckName)
		if err != nil {
			session = sessions.NewSession(store, ckName)
		}
		value := session.Values["test"]
		log.Printf("session に保存された値は %v です", value)
		// ここで何かしら session から取り出した値を使って認証だったりと言った処理を行う。
		next.ServeHTTP(rw, r)
	})
}
