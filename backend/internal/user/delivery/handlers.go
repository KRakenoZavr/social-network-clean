package delivery

import (
	"encoding/json"
	"log"
	"net/http"

	"mux/internal/middleware"
	"mux/internal/models"
	"mux/internal/user/dto"
	"mux/pkg/utils/errHandler"

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

		dbUser, cookie, sError := h.userUC.Create(rBody)
		if sError.Err != nil {
			h.logger.Println(sError.Error())
			sError.ErrorResponse(w)
			return
		}

		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(dbUser)
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

func (h userHandlers) Follow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rBody := &models.UserFollow{}
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			h.logger.Println(err.Error())
			errHandler.ErrorResponse(w, http.StatusBadRequest, err, []string{})
			return
		}

		user := r.Context().Value(middleware.ContextUserKey).(models.User)

		sError := h.userUC.Follow(rBody, user)
		if sError.Err != nil {
			h.logger.Println(sError.Error())
			sError.ErrorResponse(w)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func (h userHandlers) GetFollows() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(middleware.ContextUserKey).(models.User)

		follows, sError := h.userUC.GetFollow(user)
		if sError.Err != nil {
			h.logger.Println(sError.Error())
			sError.ErrorResponse(w)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(follows)
	}
}

func (h *userHandlers) Resolve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rBody := &dto.ModelResolve{}
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			h.logger.Println(err.Error())
			errHandler.ErrorResponse(w, http.StatusBadRequest, err, []string{})
			return
		}

		user := r.Context().Value(middleware.ContextUserKey).(models.User)

		sError := h.userUC.Resolve(rBody, user)
		if sError.Err != nil {
			h.logger.Println(sError.Error())
			sError.ErrorResponse(w)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
