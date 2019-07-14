package main

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fharding1/gemux"

	"emahiro/il/gae_sandbox/model"
)

var addr = 8080

func main() {
	mux := &gemux.ServeMux{}
	mux.Handle("/hello", http.MethodGet, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	}))
	mux.Handle("/metadata", http.MethodGet, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := http.Client{Transport: http.DefaultTransport}

		values := url.Values{}
		values.Set("audience", "test")
		values.Set("format", "full")

		reqURL := fmt.Sprintf("http://metadata/computeMetadata/v1/instance/service-accounts/default/identity?%s", values.Encode())
		req, err := http.NewRequest(http.MethodGet, reqURL, nil)
		if err != nil {
			log.Printf("failed to create request")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		req.Header.Set("Metadata-Flavor", "Google")

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("failed to get metadata. err: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		w.WriteHeader(resp.StatusCode)

		if resp.StatusCode != http.StatusOK {
			log.Printf("failed to get metadata. body: %+v", resp.Body)
			return
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("failed to read body. err: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Write(b)
	}))
	mux.Handle("/verify", http.MethodGet, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := http.Client{Transport: http.DefaultTransport}
		req, err := http.NewRequest(http.MethodGet, "https://www.googleapis.com/oauth2/v3/certs", nil)
		if err != nil {
			log.Printf("failed to create request. err: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("failed to get public keys. err: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		w.WriteHeader(resp.StatusCode)

		if resp.StatusCode != http.StatusOK {
			log.Printf("failed to get public kyes. code: %d", resp.StatusCode)
			return
		}

		keys := model.PublicKeys{}
		if err := json.NewDecoder(resp.Body).Decode(&keys); err != nil {
			log.Printf("failed to decode json. err: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// verify
		hdr := r.Header.Get("Authorization")
		if hdr == "" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		p := strings.Split(hdr, " ")
		if len(p) != 2 {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		if p[0] != "Bearer" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		parsedToken, err := jwt.Parse(p[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				err := errors.New("unexpected signing method")
				return nil, err
			}

			k := keys.Keys[1]

			dn, err := base64.RawURLEncoding.DecodeString(k.N)
			if err != nil {
				return nil, err
			}
			n := new(big.Int).SetBytes(dn)
			log.Printf("n: %d", n)

			de, err := base64.RawURLEncoding.DecodeString(k.E)
			if err != nil {
				return nil, err
			}
			e := binary.BigEndian.Uint16(de)
			log.Printf("e: %v", e)

			return &rsa.PublicKey{
				N: n,
				E: int(e),
			}, nil
		})
		if err != nil {
			log.Printf("failed to parse token. err: %v", err)
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		if !parsedToken.Valid {
			log.Printf("failed to validation token. token: %+v", parsedToken)
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		log.Printf("success velity token")
		w.WriteHeader(http.StatusOK)

	}))

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", addr),
		Handler: mux,
	}

	log.Printf("server start port: %d ...", addr)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("server closed with error: %v", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	log.Printf("SIGNAL %d received, server shutdown ...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("failed to gracefully shutdown. err: %v", err)
		os.Exit(1)
	}
	log.Printf("server shutdown....")
}
