package internal

import (
	"database/sql"
	"fmt"
	"net/http"

	"mux/pkg/logger"

	groupHttp "mux/internal/group/delivery"
	groupRepository "mux/internal/group/repository"
	groupUseCase "mux/internal/group/usecase"
	"mux/internal/middleware"
	userHttp "mux/internal/user/delivery"
	userRepository "mux/internal/user/repository"
	userUseCase "mux/internal/user/usecase"
	"mux/pkg/db/sqlite"
	"mux/pkg/server/controller"
	"mux/pkg/server/router"
	"mux/pkg/utils"
)

type Server struct {
	router *router.Router
	db     *sql.DB
}

func NewServer() (s *Server) {
	s = &Server{
		router: router.NewRouter(),
		db:     sqlite.CreateDB(true),
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
		Handler: c.Logging(s.other(s.router)),
	}

	return server.ListenAndServe()
}

func (s *Server) configureRouter() {
	handlersLogger := logger.HandlersLogger()

	// init repo
	userRepo := userRepository.NewRepository(s.db, handlersLogger)
	groupRepo := groupRepository.NewGroupRepository(s.db, handlersLogger)

	// init usecase
	userUC := userUseCase.NewUseCase(userRepo, handlersLogger)
	groupUC := groupUseCase.NewGroupUseCase(groupRepo, handlersLogger)

	// init handler
	userHandlers := userHttp.NewHandler(userUC, handlersLogger)
	groupHandlers := groupHttp.NewGroupHandlers(groupUC, handlersLogger)

	// init middleware
	authMW := middleware.NewAuthMiddleware(userRepo, handlersLogger)

	userHttp.MapRoutes(s.router, userHandlers)
	groupHttp.MapGroupRoutes(s.router, groupHandlers)

	s.router.HandleFunc("/", authMW.CheckAuth(asd())).Methods("GET")
}

func asd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}
}

func (s *Server) other(hdlr http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pathes := s.router.GetPathes()
		if !utils.Contains(pathes, r.URL.Path) {
			w.WriteHeader(http.StatusNotFound)
		}

		hdlr.ServeHTTP(w, r)
	})
}
