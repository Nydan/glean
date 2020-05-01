package app

import (
	"github.com/nydan/glean/internal/config"
	"github.com/nydan/glean/internal/controller/http"
	"github.com/nydan/glean/internal/server"
)

// HTTPServer initializes all dependencies for HTTP server
func HTTPServer(cfg config.Config) error {

	ctrl := http.NewController()

	srv := server.NewHTTPServer(cfg.HTTPServer, ctrl)

	return srv.RunHTTP()
}
