package errHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrResponse struct {
	Code 	int
	Message []string
	Err 	string
}

type ServiceError struct {
	Code    int
	Message []string
	Err     error
}

func (s *ServiceError) Error() string {
	return fmt.Sprintf("got error with status: %v, err: %v, msg: %s", s.Code, s.Err, s.Message)
}

func (s *ServiceError) ErrorResponse(w http.ResponseWriter) {
	w.WriteHeader(s.Code)
	json.NewEncoder(w).Encode(s.createResponse())
}

func (s *ServiceError) createResponse() *ErrResponse {
	return &ErrResponse{Code: s.Code, Message: s.Message, Err: s.Err.Error()}
}

func ErrorResponse(w http.ResponseWriter, code int, err error, msg []string) {
	sError := &ServiceError{Code: code, Message: msg, Err: err}
	sError.ErrorResponse(w)
}
