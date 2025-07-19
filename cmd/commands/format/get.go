package format

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"simple-cli/internal/config"

	"github.com/spf13/cobra"
)

type GetFormatCommand struct {
	cfg    *config.Config
	logger *slog.Logger
	Cmd    *cobra.Command
}

func NewGetFormatCommand(cfg *config.Config, logger *slog.Logger) *GetFormatCommand {
	getForGetFormatCommand := &GetFormatCommand{
		cfg:    cfg,
		logger: logger,
		Cmd: &cobra.Command{
			Use:   "get",
			Short: "Show current file format",
			Run: func(cmd *cobra.Command, args []string) {
				getFormat(os.Stdout, cfg)
			},
		},
	}

	return getForGetFormatCommand
}

// getFormat выводит формат файлов (экспортируемая для тестов)
func getFormat(out io.Writer, cfg *config.Config) {
	fmt.Fprintf(out, "%s", cfg.File.Format)
}
