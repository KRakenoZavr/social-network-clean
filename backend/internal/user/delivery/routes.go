package delivery

import (
	"mux/internal/middleware"
	"mux/internal/user"
	"mux/pkg/server/router"
)

func MapRoutes(r *router.Router, h user.Handlers, mw *middleware.AuthMiddleware) {
	// user registration route
	r.HandleFunc("/user/register", h.Create()).Methods("POST")
	// user login
	r.HandleFunc("/user/login", h.Login()).Methods("POST")
	// user request to join group
	r.HandleFunc("/user/follow", mw.CheckAuth(h.Follow())).Methods("POST")
}
