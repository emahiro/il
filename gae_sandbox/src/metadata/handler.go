package metadata

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat/go-jwx/jwk"
)

func GetMetadata(w http.ResponseWriter, r *http.Request) {
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
}

func Verify(w http.ResponseWriter, r *http.Request) {
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

	token, err := jwt.Parse(p[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			err := errors.New("unexpected signing method")
			return nil, err
		}

		start := time.Now()
		set, err := jwk.Fetch("https://www.googleapis.com/oauth2/v3/certs")
		if err != nil {
			log.Printf("cannot get fetch key set. err: %v", err)
			return nil, err
		}
		end := time.Now()
		d := end.Sub(start)
		log.Printf("duration: %v", d)

		keyID, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("expecting JWT header to have string kid")
		}

		key := set.LookupKeyID(keyID)
		if len(key) != 1 {
			return nil, fmt.Errorf("unable to find key")
		}

		return key[0].Materialize()
	})
	if err != nil {
		log.Printf("failed to parse token. err: %v", err)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}
	if !token.Valid {
		log.Printf("failed to validation token. token: %+v", token)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	log.Printf("success velity token")
	log.Printf("%+v", token)
	w.WriteHeader(http.StatusOK)
}
