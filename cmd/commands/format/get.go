package format

import (
	"fmt"
	"log/slog"
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
				getFormat(cfg)
			},
		},
	}

	return getForGetFormatCommand
}

func getFormat(cfg *config.Config) {
	fmt.Println(cfg.File.Format)
}
