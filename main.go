package main

import (
	"simple-cli/cmd"
	"simple-cli/internal/config"
	"simple-cli/internal/logger"
)

func main() {
	cfg := config.MustLoad()

	logger := logger.SetupLogger(cfg.Log.Level)

	app := cmd.NewApp(cfg, logger)
	app.Execute()
}