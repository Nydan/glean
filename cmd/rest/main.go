package main

import (
	"flag"

	"github.com/nydan/glean/internal/app"
	"github.com/nydan/glean/internal/config"
	"github.com/nydan/glean/internal/environment"
	"github.com/nydan/glean/pkg/logger"
)

func main() {
	env := flag.String("env", environment.Local, "Environment of the app (local, integration, production). Default is local.")
	flag.Parse()

	validateEnv(*env)

	cfg, err := config.Load(*env)
	if err != nil {
		panic("Failed to read config: " + err.Error())
	}

	logger.NewLogger((logger.Configuration)(cfg.Logger))

	err = app.HTTPServer(*cfg)
	if err != nil {
		logger.Fatalw("Stop serving", "error", err)
	}
}

// validateEnv validates the environment from flags
func validateEnv(env string) {
	if env != environment.Local && env != environment.Integration && env != environment.Production {
		panic("Unknown environment : " + env)
	}
}
