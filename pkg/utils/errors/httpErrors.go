package errors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

var (
	ErrMethodMismatch = errors.New("method is not allowed")
	ErrNotFound       = errors.New("no matching route was found")
)

type ErrResponse struct {
	Code    int
	Message string
}

func ErrorResponse(logger *log.Logger, w http.ResponseWriter, code int, err error) {
	response := ErrResponse{Code: code, Message: err.Error()}
	logger.Println(err.Error())
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
