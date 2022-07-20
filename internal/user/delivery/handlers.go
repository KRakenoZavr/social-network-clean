package delivery

import (
	"encoding/json"
	"log"
	"mux/internal/models"
	"mux/pkg/utils/errHandler"
	"net/http"

	"mux/internal/user"
)

type userHandlers struct {
	userUC user.UseCase
	logger *log.Logger
}

func NewHandler(userUC user.UseCase, logger *log.Logger) user.Handlers {
	return &userHandlers{userUC: userUC, logger: logger}
}

func (h userHandlers) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rBody := &models.User{}
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			h.logger.Println(err.Error())
			errHandler.ErrorResponse(w, http.StatusBadRequest, err, []string{})
			return
		}

		cookie, sError := h.userUC.Create(rBody)
		if sError.Err != nil {
			h.logger.Println(sError.Error())
			sError.ErrorResponse(w)
			return
		}

		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusCreated)
	}
}

func (h userHandlers) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rBody := &models.UserLogin{}
		err := json.NewDecoder(r.Body).Decode(rBody)
		if err != nil {
			h.logger.Println(err.Error())
			errHandler.ErrorResponse(w, http.StatusBadRequest, err, []string{})
			return
		}

		cookie, sError := h.userUC.Login(rBody)
		if sError.Err != nil {
			h.logger.Println(sError.Error())
			sError.ErrorResponse(w)
			return
		}

		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusOK)
	}
}
