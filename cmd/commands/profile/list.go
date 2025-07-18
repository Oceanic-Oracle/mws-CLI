package profile

import (
	"log/slog"

	"github.com/spf13/cobra"
)

type ListProfileCommand struct {
	logger *slog.Logger
	Cmd    *cobra.Command
}

func NewListProfileCommand(logger *slog.Logger) *ListProfileCommand {
	listProfileCommand := &ListProfileCommand{
		logger: logger,
		Cmd: &cobra.Command{
			Use:   "list",
			Short: "List all available profiles",
			Run: func(cmd *cobra.Command, args []string) {
				listProfile()
			},
		},
	}

	return listProfileCommand
}

func listProfile() {
}
