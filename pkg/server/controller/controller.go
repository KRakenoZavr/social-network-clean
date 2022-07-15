package controller

import (
	"log"
	"net/http"
	"os"
	"time"

	colors "mux/pkg/server/colors"
)

func NewController() *Controller {
	logger := log.New(os.Stdout, colors.CLIBlue, log.LstdFlags)

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
