package middleware

import (
	"errors"
	"log"
	"mux/pkg/utils/errHandler"
	"net/http"
)

type (
	Middleware  func(http.Handler) http.Handler
	Middlewares []Middleware
)

func checkAuth(c *http.Cookie) error {
	c.Expires.Add(5 * 60 * 1000)
	if false {
		return errors.New("not authed")
	}
	return nil
}

func CheckAuth(hdlr http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("secret")
		if err != nil {
			log.Println("error accessing cookie", err.Error())
			errHandler.ErrorResponse(w, http.StatusInternalServerError, errors.New("error accessing cookie"), []string{})
			return
		}

		err = checkAuth(cookie)
		if err != nil {
			errHandler.ErrorResponse(w, http.StatusUnauthorized, errors.New("you should authorize"), []string{})
			return
		}

		hdlr.ServeHTTP(w, r)

	})
}
