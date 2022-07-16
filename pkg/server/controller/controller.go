package controller

import (
	"log"
	"mux/pkg/colors"
	logger2 "mux/pkg/logger"
	"net/http"
	"time"
)

func NewController() *Controller {
	logger := logger2.ServerLogger()

	return &Controller{Logger: logger}
}

type Controller struct {
	Logger *log.Logger
}

func (c *Controller) Logging(hdlr http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func(start time.Time) {
			c.Logger.Println(colors.CLIClear, r.Method, r.URL.Path, time.Since(start))
		}(time.Now())
		hdlr.ServeHTTP(w, r)
	})
}
