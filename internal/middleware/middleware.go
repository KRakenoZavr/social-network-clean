package middleware

import (
	"errors"
	"log"
	errorHandler "mux/pkg/utils/errors"
	"net/http"
)

type (
	Middleware  func(http.Handler) http.Handler
	Middlewares []Middleware
)

func checkAuth(c *http.Cookie) error {
	c.Expires.Add(5 * 60 * 1000)
	if false {
		return errors.New("asd")
	}
	return nil
}

func CheckAuth(hdlr http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("secret")
		if err != nil {
			log.Println("error accessing cookie", err.Error())
			errorHandler.ErrorResponse(w, http.StatusInternalServerError, errors.New("error accessing cookie"))
			return
		}

		err = checkAuth(cookie)
		if err != nil {
			errorHandler.ErrorResponse(w, http.StatusUnauthorized, errors.New("you should authorize"))
			return
		}

		hdlr.ServeHTTP(w, r)

	})
}
