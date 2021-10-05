package api

import (
	"go.uber.org/zap"
	"net/http"
	"github.com/gorilla/mux"
	"time"
)

type Server struct {
	router  *mux.Router
	logger  *zap.Logger
	handler http.Handler
}

func NewServer(logger *zap.Logger) (*Server, error) {
	srv := &Server{
		router: mux.NewRouter(),
		logger: logger,
	}

	return srv, nil
}

func (s *Server) registerHandlers() {
	s.router.HandleFunc("/token", s.tokenGenerateHandler).Methods("POST")
	s.router.HandleFunc("/public-key", s.publicKeyHandler).Methods("GET")

	httpLogger := NewLoggingMiddleware(s.logger)
	s.router.Use(httpLogger.Handler)
}

func (s *Server) ListenAndServe() {
	s.registerHandlers()

	srv := &http.Server{
		Handler: s.router,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	s.logger.Info("Starting HTTP Server.", zap.String("addr", srv.Addr))
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		s.logger.Fatal("HTTP server crashed", zap.Error(err))
	}
}