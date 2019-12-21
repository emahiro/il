package mw

import (
	"fmt"
	"net/http"

	"github.com/emahiro/log_output/logger"
)

func Logger() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			line := fmt.Sprintf("%s %s", r.Method, r.URL.Path)
			logger.Debugf("%s", line)
			h.ServeHTTP(w, r)
		})
	}
}
