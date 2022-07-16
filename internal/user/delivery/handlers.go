package delivery

import (
	"log"
	"net/http"

	"mux/internal/user"
)

type userHandlers struct {
	userUC user.UseCase
	logger *log.Logger
}

func NewUserHandlers(u user.UseCase, l *log.Logger) user.Handlers {
	return &userHandlers{userUC: u, logger: l}
}

func (h userHandlers) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.logger.Println("asd")
	}
}
