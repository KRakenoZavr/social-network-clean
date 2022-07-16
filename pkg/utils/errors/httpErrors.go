package errors

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrResponse struct {
	Code int
	Message string
}

func ErrorResponse(logger *log.Logger, w http.ResponseWriter, code int, err error) {
	response := ErrResponse{Code: code, Message: err.Error()}
	logger.Println(err.Error())
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
