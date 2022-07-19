package middleware

import (
	"errors"
	"log"
	"mux/internal/user"
	"mux/pkg/utils/errHandler"
	"net/http"
	"time"
)

type (
	Middleware  func(http.Handler) http.Handler
	Middlewares []Middleware
)

var (
	TokenExpired = errors.New("session token expired")
)

type AuthMiddleware struct {
	r      user.Repository
	logger *log.Logger
}

func NewAuthMiddleware(r user.Repository, logger *log.Logger) *AuthMiddleware {
	return &AuthMiddleware{r: r, logger: logger}
}

func (m *AuthMiddleware) checkAuth(c *http.Cookie) error {
	userAuth, err := m.r.GetUserAuth(c.Value)
	if err != nil {
		return err
	}

	now := time.Now()
	isBefore := userAuth.Expires.Before(now)

	if isBefore {
		return TokenExpired
	}
	return nil
}

func (m *AuthMiddleware) CheckAuth(hdlr http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			log.Println("no cookie", err.Error())
			errHandler.ErrorResponse(w, http.StatusUnauthorized, err, []string{"no cookie"})
			return
		}

		err = m.checkAuth(cookie)
		if err != nil {
			errHandler.ErrorResponse(w, http.StatusUnauthorized, err, []string{})
			return
		}

		hdlr(w, r)
	})
}
