package group

import "net/http"

type Handlers interface {
	Create() http.HandlerFunc
	JoinRequest() http.HandlerFunc
	Get() http.HandlerFunc
	GetRequests() http.HandlerFunc
	Invite() http.HandlerFunc
	GetInvites() http.HandlerFunc
	Resolve() http.HandlerFunc
}
