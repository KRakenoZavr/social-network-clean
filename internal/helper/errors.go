package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrResponse struct {
	ErrorText string
	Status    int
}

func ResponseWithErr(w http.ResponseWriter, status int, errText string) {
	log.Println(errText)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrResponse{ErrorText: errText, Status: status})
}
