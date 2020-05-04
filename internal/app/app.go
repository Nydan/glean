package app

import (
	"github.com/nydan/glean/internal/config"
	"github.com/nydan/glean/internal/controller/http"
	ordctrl "github.com/nydan/glean/internal/controller/http/order"
	"github.com/nydan/glean/internal/server"
	orduc "github.com/nydan/glean/internal/usecase/order"
)

// HTTPServer initializes all dependencies for HTTP server.
// The initialization that happen here are related to HTTP API.
func HTTPServer(cfg config.Config) error {
	orderUc := orduc.NewOrder(struct{}{})

	orderCtrl := ordctrl.Order(orderUc)

	ctrl := http.NewController(orderCtrl)

	srv := server.NewHTTPServer(cfg.HTTPServer, ctrl)

	return srv.RunHTTP()
}
