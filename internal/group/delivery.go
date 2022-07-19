package group

import "net/http"

type Handlers interface {
	Create() http.HandlerFunc
}
