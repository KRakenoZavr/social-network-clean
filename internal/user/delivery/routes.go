package delivery

import (
	"mux/internal/user"
	"mux/pkg/server/router"
)

func MapUserRoutes(r *router.Router, h user.Handlers) {
	r.HandleFunc("/user/register", h.Create()).Methods("POST")
	r.HandleFunc("/user/login", h.Login()).Methods("POST")
}
