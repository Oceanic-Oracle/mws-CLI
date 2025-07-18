package profile

import (
	"log/slog"

	"github.com/spf13/cobra"
)

type GetProfileCommand struct {
	logger *slog.Logger
	Cmd    *cobra.Command
}

func NewGetProfileCommand(logger *slog.Logger) *GetProfileCommand {
	getProfileCommand := &GetProfileCommand{
		logger: logger,
		Cmd: &cobra.Command{
			Use:   "get",
			Short: "Display details",
			Run: func(cmd *cobra.Command, args []string) {
				getProfile()
			},
		},
	}

	return getProfileCommand
}

func getProfile() {
}
