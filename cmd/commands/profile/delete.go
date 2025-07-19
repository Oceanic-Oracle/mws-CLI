package profile

import (
	"fmt"
	"log/slog"
	"os"
	profile_storage "simple-cli/internal/storage/profile"

	"github.com/spf13/cobra"
)

type DeleteProfileCommand struct {
	logger *slog.Logger
	Cmd    *cobra.Command
}

func NewDeleteProfileCommand(logger *slog.Logger, profileStorage profile_storage.IProfile) *DeleteProfileCommand {
	deleteProfileCommand := &DeleteProfileCommand{
		logger: logger,
		Cmd: &cobra.Command{
			Use:   "delete",
			Short: "Delete an existing profile",
			Run: func(cmd *cobra.Command, args []string) {
				name, _ := cmd.Flags().GetString("name")

				var path string
				if len(args) > 0 {
					path = args[0]
				}

				deleteProfile(profileStorage, path, name)
			},
		},
	}

	deleteProfileCommand.Cmd.Flags().StringP("name", "", "", "Profile name (required)")
	deleteProfileCommand.Cmd.MarkFlagRequired("name")

	return deleteProfileCommand
}

// deleteProfile удаляет профиль (экспортируемая для тестов)
func deleteProfile(profileStorage profile_storage.IProfile, path, name string) {
	if err := profileStorage.Delete(path, name); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
