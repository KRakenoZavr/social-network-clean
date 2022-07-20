package delivery

import (
	"encoding/json"
	"log"
	"mux/internal/models"
	"mux/pkg/utils/errHandler"
	"net/http"

	"mux/internal/group"
)

type groupHandlers struct {
	groupUC group.UseCase
	logger  *log.Logger
}

func NewGroupHandlers(groupUC group.UseCase, logger *log.Logger) group.Handlers {
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

		sError := h.groupUC.Create(rBody)
		if sError.Err != nil {
			h.logger.Println(sError.Error())
			sError.ErrorResponse(w)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
