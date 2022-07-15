package middleware

import "net/http"

type (
	Middleware  func(http.Handler) http.Handler
	Middlewares []Middleware
)
