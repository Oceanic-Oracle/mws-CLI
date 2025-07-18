package profile

import (
	"log/slog"

	"github.com/spf13/cobra"
)

type DeleteProfileCommand struct {
	logger *slog.Logger
	Cmd    *cobra.Command
}

func NewDeleteProfileCommand(logger *slog.Logger) *DeleteProfileCommand {
	deleteProfileCommand := &DeleteProfileCommand{
		logger: logger,
		Cmd: &cobra.Command{
			Use:   "delete",
			Short: "Delete an existing profile",
			Run: func(cmd *cobra.Command, args []string) {
				deleteProfile()
			},
		},
	}

	return deleteProfileCommand
}

func deleteProfile() {
	
}
