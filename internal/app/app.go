package app

import (
	"github.com/nydan/glean/internal/config"
	"github.com/nydan/glean/internal/controller/http"
	"github.com/nydan/glean/internal/controller/http/order"
	"github.com/nydan/glean/internal/server"
)

// HTTPServer initializes all dependencies for HTTP server.
// The initialization that happen here are related to HTTP API.
func HTTPServer(cfg config.Config) error {
	orderUc := order.Controller{}
	ctrl := http.NewController(orderUc)

	srv := server.NewHTTPServer(cfg.HTTPServer, ctrl)

	return srv.RunHTTP()
}
