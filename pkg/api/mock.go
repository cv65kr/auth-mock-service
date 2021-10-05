package api

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func NewMockServer() *Server {
	logger, _ := zap.NewDevelopment()

	return &Server{
		router: mux.NewRouter(),
		logger: logger,
	}
}