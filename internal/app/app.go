package app

import (
	"github.com/nydan/glean/internal/config"
	"github.com/nydan/glean/internal/server"
)

// HTTPServer initializes all dependencies for HTTP server
func HTTPServer(cfg config.Config) error {
	srv := server.NewHTTPServer(cfg.HTTPServer)

	return srv.RunHTTP()
}
