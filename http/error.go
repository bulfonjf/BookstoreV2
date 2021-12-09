package http

import (
	"bookstore/application"
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func handleErrorAsJson(w http.ResponseWriter, r *http.Request, code int, message string, err error) {
	if err != nil {
		logError(r, err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&ErrorResponse{Error: message})

}

var codes = map[string]int{
	application.ECONFLICT:       http.StatusConflict,
	application.EINVALID:        http.StatusBadRequest,
	application.ENOTFOUND:       http.StatusNotFound,
	application.ENOTIMPLEMENTED: http.StatusNotImplemented,
	application.EUNAUTHORIZED:   http.StatusUnauthorized,
	application.EINTERNAL:       http.StatusInternalServerError,
}

func errorStatusCode(code string) int {
	if v, ok := codes[code]; ok {
		return v
	}
	return http.StatusInternalServerError
}

func logError(r *http.Request, err error) {
	log.Printf("[http] error: %s %s: %s", r.Method, r.URL.Path, err)
}
