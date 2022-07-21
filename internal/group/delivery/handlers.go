package delivery

import (
	"encoding/json"
	"log"
	"net/http"

	"mux/internal/middleware"
	"mux/internal/models"
	"mux/pkg/utils/errHandler"

	"mux/internal/group"
)

type groupHandlers struct {
	groupUC group.UseCase
	logger  *log.Logger
}

func NewHandler(groupUC group.UseCase, logger *log.Logger) group.Handlers {
	return &groupHandlers{groupUC: groupUC, logger: logger}
}

func (h groupHandlers) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rBody := &models.Group{}
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			h.logger.Println(err.Error())
			errHandler.ErrorResponse(w, http.StatusBadRequest, err, []string{})
			return
		}

		user := r.Context().Value(middleware.ContextUserKey).(models.User)

		sError := h.groupUC.Create(rBody, user)
		if sError.Err != nil {
			h.logger.Println(sError.Error())
			sError.ErrorResponse(w)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func (h *groupHandlers) JoinRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rBody := &models.GroupUser{}
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			h.logger.Println(err.Error())
			errHandler.ErrorResponse(w, http.StatusBadRequest, err, []string{})
			return
		}

		user := r.Context().Value(middleware.ContextUserKey).(models.User)

		sError := h.groupUC.JoinRequest(rBody, user)
		if sError.Err != nil {
			h.logger.Println(sError.Error())
			sError.ErrorResponse(w)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func (h *groupHandlers) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groups, sError := h.groupUC.GetAllGroups()
		if sError.Err != nil {
			h.logger.Println(sError.Error())
			sError.ErrorResponse(w)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(groups)
	}
}

func (h *groupHandlers) GetRequests() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(middleware.ContextUserKey).(models.User)

		gUsers, sError := h.groupUC.GetRequests(user)
		if sError.Err != nil {
			h.logger.Println(sError.Error())
			sError.ErrorResponse(w)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(gUsers)
	}
}

func (h *groupHandlers) Invite() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rBody := &models.GroupUser{}
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			h.logger.Println(err.Error())
			errHandler.ErrorResponse(w, http.StatusBadRequest, err, []string{})
			return
		}

		user := r.Context().Value(middleware.ContextUserKey).(models.User)

		sError := h.groupUC.Invite(rBody, user)
		if sError.Err != nil {
			h.logger.Println(sError.Error())
			sError.ErrorResponse(w)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func (h *groupHandlers) GetInvites() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(middleware.ContextUserKey).(models.User)

		groups, sError := h.groupUC.GetInvites(user)
		if sError.Err != nil {
			h.logger.Println(sError.Error())
			sError.ErrorResponse(w)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(groups)
	}
}
