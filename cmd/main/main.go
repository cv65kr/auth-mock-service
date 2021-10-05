package main

import (
	"auth-mock-service/pkg/api"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	stdLog := zap.RedirectStdLog(logger)
	defer stdLog()

	// start HTTP server
	srv, _ := api.NewServer(logger)
	srv.ListenAndServe()
}
