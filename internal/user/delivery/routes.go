package delivery

import (
	"mux/internal/user"
	"mux/pkg/server/router"
)

func MapUserRoutes(r *router.Router, h user.Handlers) {
	r.HandleFunc("/user/create", h.Create()).Methods("POST")
}
