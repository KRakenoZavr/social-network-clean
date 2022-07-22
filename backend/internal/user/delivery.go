package user

import "net/http"

type Handlers interface {
	Create() http.HandlerFunc
	Login() http.HandlerFunc
	Follow() http.HandlerFunc
	GetFollows() http.HandlerFunc
	Resolve() http.HandlerFunc
}
