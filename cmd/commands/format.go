package commands

import (
	"log/slog"
	"simple-cli/cmd/commands/format"
	"simple-cli/internal/config"

	"github.com/spf13/cobra"
)

type FormatCommand struct {
	Cmd    *cobra.Command
}

func NewFormatCommand(cfg *config.Config, logger *slog.Logger) *FormatCommand {
	formatCommand := &FormatCommand{
		Cmd: &cobra.Command{
			Use:   "format",
			Short: "Change format files",
		},
	}

	formatCommand.Cmd.AddCommand(format.NewGetFormatCommand(cfg, logger).Cmd,
		format.NewSetFormatCommand(cfg, logger).Cmd)

	return formatCommand
}
