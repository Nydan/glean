package main

import (
	"flag"

	"github.com/nydan/glean/internal/app"
	"github.com/nydan/glean/internal/config"
	"github.com/nydan/glean/internal/environment"
	"github.com/nydan/glean/pkg/slog"
)

func main() {
	env := flag.String("env", environment.Local, "Environment of the app (local, integration, production). Default is local.")
	flag.Parse()

	validateEnv(*env)

	cfg, err := config.Load(*env)
	if err != nil {
		panic("Failed to read config: " + err.Error())
	}

	slog.NewLogger((slog.Configuration)(cfg.Logger))

	err = app.HTTPServer(*cfg)
	if err != nil {
		slog.Fatalw("Stop serving", "error", err)
	}
}

// validateEnv validates the environment from flags
func validateEnv(env string) {
	if env != environment.Local && env != environment.Integration && env != environment.Production {
		panic("Unknown environment : " + env)
	}
}
