package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"simple-cli/cmd/commands"
	"simple-cli/internal/config"

	"github.com/spf13/cobra"
)

type App struct {
	cfg *config.Config
	logger *slog.Logger
	rootCmd *cobra.Command
}

func NewApp(cfg *config.Config, logger *slog.Logger) *App {
	app := &App{
		cfg: cfg,
		logger: logger,
		rootCmd: &cobra.Command{
			Use:   "mws",
			Short: "Profile manager",
		},
	}

	app.rootCmd.AddCommand(commands.NewFormatCommand(app.cfg, app.logger).Cmd,
		commands.NewProfileCommand(app.cfg, app.logger).Cmd)

	return app
}

func (app *App) Execute() {
	if err := app.rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
