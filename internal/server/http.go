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
	"github.com/nydan/glean/pkg/slog"
)

// HTTPServerI is interface to wrap http server library
type HTTPServerI interface {
	RunHTTP() error
}

type httpServer struct {
	srv http.Server
}

// NewHTTPServer creates new http server
func NewHTTPServer(cfg config.HTTPServer) HTTPServerI {
	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.ListenAddress, cfg.Port),
		Handler: router(),
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
			slog.Errorw("HTTP Server shutting down", "error", err)
		}

		close(done)
	}()

	slog.Infow("HTTP server running on ", "address", h.srv.Addr)
	if err := h.srv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	<-done

	slog.Infow("HTTP server shutdown gracefully")
	return nil
}

func router() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	}))
	return router
}
