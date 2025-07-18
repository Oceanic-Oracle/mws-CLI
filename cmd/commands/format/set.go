package format

import (
	"fmt"
	"log/slog"
	"simple-cli/internal/config"

	"github.com/spf13/cobra"
)

type SetFormatCommand struct {
	cfg    *config.Config
	logger *slog.Logger
	Cmd    *cobra.Command
}

func NewSetFormatCommand(cfg *config.Config, logger *slog.Logger) *SetFormatCommand {
	setFormatCommand := &SetFormatCommand{
		cfg:    cfg,
		logger: logger,
		Cmd: &cobra.Command{
			Use:   "set",
			Short: "Change file format",
			Run: func(cmd *cobra.Command, args []string) {
				format, _ := cmd.Flags().GetString("format")
				setFormat(cfg, format)
			},
		},
	}

	setFormatCommand.Cmd.Flags().StringP("format", "f", "yaml", "File format")

	return setFormatCommand
}

func setFormat(cfg *config.Config, format string) {
	if cfg.IsDefault {
		fmt.Println("Note: Using default configuration")
	}
	
	cfg.File.Format = format
	cfg.Save()
}
