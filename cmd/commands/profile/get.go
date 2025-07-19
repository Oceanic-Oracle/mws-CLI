package profile

import (
	"fmt"
	"io"
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

				GetProfile(os.Stdout, profileStorage, path, name)
			},
		},
	}

	getProfileCommand.Cmd.Flags().StringP("name", "", "", "Profile name (required)")
	getProfileCommand.Cmd.MarkFlagRequired("name")

	return getProfileCommand
}

// GetProfile выводит содержимое профиля (экспортируемая для тестов)
func GetProfile(out io.Writer, profileStorage profile_storage.IProfile, path, name string) {
	data, err := profileStorage.Get(path, name)
	if err != nil {
		fmt.Fprintf(out, "%v", err)
		return
	}

	fmt.Fprintf(out, "%s", data)
}
