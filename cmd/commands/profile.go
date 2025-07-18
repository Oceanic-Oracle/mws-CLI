package commands

import (
	"log/slog"
	"simple-cli/cmd/commands/profile"
	"simple-cli/internal/config"
	profile_storage "simple-cli/internal/storage/profile"
	profile_yaml "simple-cli/internal/storage/profile/yaml"

	"github.com/spf13/cobra"
)

type ProfileCommand struct {
	Cmd     *cobra.Command
	Storage profile_storage.IProfile
}

func NewProfileCommand(cfg *config.Config, logger *slog.Logger) *ProfileCommand {
	profileCommand := &ProfileCommand{
		Cmd: &cobra.Command{
			Use:   "profile",
			Short: "Profile Management Commands",
		},
		Storage: profile_yaml.NewProfileStorage(cfg, logger),
	}

	profileCommand.Cmd.AddCommand(
		profile.NewCreateProfileCommand(logger, profileCommand.Storage).Cmd,
		profile.NewDeleteProfileCommand(logger).Cmd,
		profile.NewGetProfileCommand(logger).Cmd,
		profile.NewListProfileCommand(logger).Cmd,
	)

	return profileCommand
}
