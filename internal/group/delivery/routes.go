package delivery

import (
	"mux/internal/group"
	"mux/pkg/server/router"
)

func MapGroupRoutes(r *router.Router, h group.Handlers) {
	r.HandleFunc("/group/create", h.Create()).Methods("POST")
}
