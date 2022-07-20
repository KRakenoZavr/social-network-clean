package delivery

import (
	"mux/internal/group"
	"mux/pkg/server/router"
)

func MapRoutes(r *router.Router, h group.Handlers) {
	r.HandleFunc("/group/create", h.Create()).Methods("POST")
}
