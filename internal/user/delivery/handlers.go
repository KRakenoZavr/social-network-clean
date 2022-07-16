package delivery

import (
	"log"
	"net/http"

	"mux/internal/user"
)

type userHandlers struct {
	userUC user.UseCase
}

func NewUserHandlers(u user.UseCase) user.Handlers {
	return &userHandlers{userUC: u}
}

func (h userHandlers) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("asd")
	}
}
