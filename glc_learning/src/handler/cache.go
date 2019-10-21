package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/emahiro/ae-plain-logger/log"
	"github.com/emahiro/glc"
)

var fc *glc.FileCache

func init() {
	var err error
	if fc, err = glc.NewFileCache("glc_learning"); err != nil {
		panic(err)
	}
}

type requestBody struct {
	Data string `json:"data"`
}

// SetCache ...
func SetCache(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var body *requestBody
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Warningf(ctx, "invalid request")
		return
	}

	b, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf(ctx, "failed to parse json. err: %v", err)
		return
	}

	if err := fc.Set("test", b); err != nil {
		log.Errorf(ctx, "failed to set error. err: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to set cache"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("set cache!"))
}

// GetCache ...
func GetCache(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := fc.Get("test")
	if data == nil {
		log.Warningf(ctx, "faild to get cache data")
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("failed to get cache"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("get cache data. data is %s", string(data))))
}
