package handler

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"emahiro/gotester/model"

	"github.com/fharding1/gemux"
)

type fakeUserB struct {
	Name string
	Age  int64
}

func (f fakeUserB) Get() (*model.UserB, error) {
	log.Printf("this is fake method")
	return &model.UserB{Name: "fakeName", Age: 999}, nil
}

func TestGetUserB(t *testing.T) {

	tests := map[string]struct {
		path string
		args struct{}
		res  struct{}
		want int64
	}{
		"case 1": {
			path: "/userB",
			args: struct{}{}, // パラメータを指定する
			res:  struct{}{}, // 結果をここでmockしたい
			want: 200,
		},
	}

	r := &gemux.ServeMux{}
	r.Handle("/userB", http.MethodGet, http.HandlerFunc(GetUserB))

	for k, tt := range tests {
		tt := tt
		t.Run(k, func(t *testing.T) {
			user := &fakeUserB{Name: "hoge", Age: 222}
			if _, err := user.Get(); err != nil {
				t.Fatalf("err: %v", err)
			}

			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080%s", tt.path), nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// b, _ := ioutil.ReadAll(w.Body)
			// t.Logf("body: %s", string(b))
		})
	}
}
