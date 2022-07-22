package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"mux/internal/models"
	"mux/internal/user"
	"mux/pkg/utils/errHandler"
)

type (
	Middleware  func(http.Handler) http.Handler
	Middlewares []Middleware
)

var TokenExpired = errors.New("session token expired")

type ContextKey string

const ContextUserKey ContextKey = "user"

type AuthMiddleware struct {
	r      user.Repository
	logger *log.Logger
}

func NewAuthMiddleware(r user.Repository, logger *log.Logger) *AuthMiddleware {
	return &AuthMiddleware{r: r, logger: logger}
}

func (m *AuthMiddleware) checkAuth(c *http.Cookie) (models.User, error) {
	userAuth, err := m.r.GetUserAuth(c.Value)
	if err != nil {
		return models.User{}, err
	}

	now := time.Now()
	isBefore := userAuth.Expires.Before(now)

	if isBefore {
		return models.User{}, TokenExpired
	}

	dbUser, err := m.r.GetUserByID(userAuth.UserID)
	if err != nil {
		return models.User{}, err
	}

	return dbUser, nil
}

func (m *AuthMiddleware) CheckAuth(hdlr http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			m.logger.Println("no cookie", err.Error())
			errHandler.ErrorResponse(w, http.StatusUnauthorized, err, []string{"no cookie"})
			return
		}

		dbUser, err := m.checkAuth(cookie)
		if err != nil {
			m.logger.Println("check cookie error", err.Error())
			errHandler.ErrorResponse(w, http.StatusUnauthorized, err, []string{})
			return
		}

		ctx := context.WithValue(r.Context(), ContextUserKey, dbUser)

		hdlr(w, r.WithContext(ctx))
	}
}
