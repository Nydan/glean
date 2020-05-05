package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nydan/glean/internal/config"
	httpctrl "github.com/nydan/glean/internal/controller/http"
	"github.com/nydan/glean/pkg/logger"
)

// HTTPServerI is interface to wrap http server library
type HTTPServerI interface {
	RunHTTP() error
}

type httpServer struct {
	srv  http.Server
	ctrl *httpctrl.Controller
}

// NewHTTPServer creates new http server
func NewHTTPServer(cfg config.HTTPServer, ctrl *httpctrl.Controller) HTTPServerI {
	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.ListenAddress, cfg.Port),
		Handler: router(ctrl),
	}
	return &httpServer{
		srv: srv,
	}
}

// RunHTTP runs http server that will gracefully shutdown on SIGTERM or SIGHUP
func (h *httpServer) RunHTTP() error {
	done := make(chan bool)
	signals := make(chan os.Signal, 1)

	// SIGHUP signals is sent when an app loses its controlling terminal
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		<-signals

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// We received an os signal, shut down.
		if err := h.srv.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			logger.Errorw("HTTP Server shutting down", "error", err)
		}

		close(done)
	}()

	logger.Infow("HTTP server running on ", "address", h.srv.Addr)
	if err := h.srv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	<-done

	logger.Infow("HTTP server shutdown gracefully")
	return nil
}
