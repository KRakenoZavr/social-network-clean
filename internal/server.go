package internal

import (
	"database/sql"
	"fmt"
	"net/http"

	"mux/pkg/logger"

	userHttp "mux/internal/user/delivery"
	userRepository "mux/internal/user/repository"
	userUseCase "mux/internal/user/usecase"
	"mux/pkg/db/sqlite"
	"mux/pkg/server/controller"
	"mux/pkg/server/router"
)

type Server struct {
	router *router.Router
	db     *sql.DB
}

func NewServer() (s *Server) {
	s = &Server{
		router: router.NewRouter(),
		db:     sqlite.CreateDB(),
	}

	s.configureRouter()

	return s
}

func (s *Server) Start(port string) error {
	s.configureRouter()

	fmt.Printf("app is running on http://localhost%s\n", port)

	c := controller.NewController()

	server := &http.Server{
		Addr:    port,
		Handler: c.Logging(s.router),
	}

	return server.ListenAndServe()
}

func (s *Server) configureRouter() {
	handlersLogger := logger.HandlersLogger()

	// init repo
	userRepo := userRepository.NewUserRepository(s.db, handlersLogger)

	// init usecase
	userUC := userUseCase.NewUserUseCase(userRepo, handlersLogger)

	// init handler
	userHandlers := userHttp.NewUserHandlers(userUC, handlersLogger)

	userHttp.MapUserRoutes(s.router, userHandlers)
}