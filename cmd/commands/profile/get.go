package profile

import (
	"fmt"
	"log/slog"
	"os"
	profile_storage "simple-cli/internal/storage/profile"

	"github.com/spf13/cobra"
)

type GetProfileCommand struct {
	logger *slog.Logger
	Cmd    *cobra.Command
}

func NewGetProfileCommand(logger *slog.Logger, profileStorage profile_storage.IProfile) *GetProfileCommand {
	getProfileCommand := &GetProfileCommand{
		logger: logger,
		Cmd: &cobra.Command{
			Use:   "get",
			Short: "Display details",
			Run: func(cmd *cobra.Command, args []string) {
				name, _ := cmd.Flags().GetString("name")

				var path string
				if len(args) > 0 {
					path = args[0]
				}

				getProfile(profileStorage, path, name)
			},
		},
	}

	getProfileCommand.Cmd.Flags().StringP("name", "", "", "Profile name (required)")
	getProfileCommand.Cmd.MarkFlagRequired("name")

	return getProfileCommand
}

// getProfile выводит содержимое профиля (экспортируемая для тестов)
func getProfile(profileStorage profile_storage.IProfile, path, name string) {
	data, err := profileStorage.Get(path, name)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Print(data)
}
