package delivery

import (
	"mux/internal/group"
	"mux/internal/middleware"
	"mux/pkg/server/router"
)

func MapRoutes(r *router.Router, h group.Handlers, mw *middleware.AuthMiddleware) {
	// create group
	r.HandleFunc("/group/create", mw.CheckAuth(h.Create())).Methods("POST")
	// user request to join group
	r.HandleFunc("/group/join", mw.CheckAuth(h.JoinRequest())).Methods("POST")
	// admin invite to group
	r.HandleFunc("/group/invite", mw.CheckAuth(h.Invite())).Methods("POST")
	// list of all groups
	r.HandleFunc("/groups", mw.CheckAuth(h.Get())).Methods("GET")
	// list of requests for admin
	r.HandleFunc("/group/check-join", mw.CheckAuth(h.GetRequests())).Methods("GET")
	// list of invites from groups
	r.HandleFunc("/group/check-invite", mw.CheckAuth(h.GetInvites())).Methods("GET")
	// accept or decline invite
	r.HandleFunc("/group/resolve", mw.CheckAuth(h.Resolve())).Methods("GET")
}
