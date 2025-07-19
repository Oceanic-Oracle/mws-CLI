package format

import (
	"fmt"
	"io"
	"log/slog"
	"os"
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
				setFormat(os.Stdout, cfg, format)
			},
		},
	}

	setFormatCommand.Cmd.Flags().StringP("format", "f", "yaml", "File format")

	return setFormatCommand
}

// setFormat устанавливает новое расширения для файлов (экспортируемая для тестов)
func setFormat(out io.Writer, cfg *config.Config, format string) {
	if cfg.IsDefault {
		fmt.Fprintf(out, "%s", "Note: Using default configuration")
		return
	}
	
	cfg.File.Format = format
	cfg.Save()
}
