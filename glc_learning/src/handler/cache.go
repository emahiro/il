package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		log.Warningf(ctx, "unexpected method")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Warningf(ctx, "failed to read body")
		return
	}

	requestBody := &requestBody{}
	if err := json.Unmarshal(body, requestBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Warningf(ctx, "invalid request. err: %v", err)
		return
	}

	b, err := json.Marshal(requestBody.Data)
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
