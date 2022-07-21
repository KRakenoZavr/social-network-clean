package delivery

import (
	"mux/internal/group"
	"mux/internal/middleware"
	"mux/pkg/server/router"
)

func MapRoutes(r *router.Router, h group.Handlers, mw *middleware.AuthMiddleware) {
	r.HandleFunc("/group/create", mw.CheckAuth(h.Create())).Methods("POST")
	r.HandleFunc("/group/join", mw.CheckAuth(h.JoinRequest())).Methods("POST")
	r.HandleFunc("/groups", mw.CheckAuth(h.Get())).Methods("GET")
}
