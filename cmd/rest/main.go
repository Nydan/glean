package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/nydan/glean/internal/app"
	"github.com/nydan/glean/internal/config"
	"github.com/nydan/glean/internal/environment"
)

func main() {
	env := flag.String("env", environment.Local, "Environment of the app (local, integration, production). Default is local.")
	flag.Parse()

	validateEnv(*env)

	cfg, err := config.Load(*env)
	if err != nil {
		panic("Failed to read config: " + err.Error())
	}

	var handler slog.Handler
	opts := &slog.HandlerOptions{Level: parseLogLevel(cfg.Logger.Level)}
	if cfg.Logger.Format == "text" {
		handler = slog.NewTextHandler(os.Stdout, opts)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}
	slog.SetDefault(slog.New(handler))

	err = app.HTTPServer(*cfg)
	if err != nil {
		slog.Error("Stop serving", "error", err)
		os.Exit(1)
	}
}

// validateEnv validates the environment from flags
func validateEnv(env string) {
	if env != environment.Local && env != environment.Integration && env != environment.Production {
		panic("Unknown environment : " + env)
	}
}

func parseLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
