package delivery

import (
	"encoding/json"
	"log"
	"mux/internal/models"
	"mux/pkg/utils/errors"
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
		rBody := &models.User{}
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			errors.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		err = h.userUC.Create(rBody)
		if err != nil {
			errors.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
