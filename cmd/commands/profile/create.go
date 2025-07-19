package profile

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	profile_storage "simple-cli/internal/storage/profile"

	"github.com/spf13/cobra"
)

type CreateProfileCommand struct {
	logger *slog.Logger
	Cmd    *cobra.Command
}

func NewCreateProfileCommand(logger *slog.Logger, profileStorage profile_storage.IProfile) *CreateProfileCommand {
	createProfileCommand := &CreateProfileCommand{
		logger: logger,
		Cmd: &cobra.Command{
			Use:   "create",
			Short: "Create a new profile",
			Args:  cobra.MaximumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				name, _ := cmd.Flags().GetString("name")
				user, _ := cmd.Flags().GetString("user")
				project, _ := cmd.Flags().GetString("project")
				
				var path string
				if len(args) > 0 {
					path = args[0]
				}

				CreateProfile(os.Stdout, profileStorage, path, name, user, project)
			},
		},
	}

	createProfileCommand.Cmd.Flags().StringP("name", "", "", "Profile name (required)")
	createProfileCommand.Cmd.Flags().StringP("user", "", "", "User name")
	createProfileCommand.Cmd.Flags().StringP("project", "", "", "Project name")
	createProfileCommand.Cmd.MarkFlagRequired("name")

	return createProfileCommand
}

// CreateProfile создает новый профиль (экспортируемая для тестов)
func CreateProfile(out io.Writer, profileStorage profile_storage.IProfile, path, name, user, project string) {
	if err := profileStorage.Create(path, name, user, project); err != nil {
		fmt.Fprintf(out, "%v", err)
		return
	}
}
