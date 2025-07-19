package commands

import (
	"log/slog"
	"simple-cli/cmd/commands/profile"
	"simple-cli/internal/config"
	profile_storage "simple-cli/internal/storage/profile"

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
		Storage: profile_storage.NewProfileStorage(cfg, logger),
	}

	profileCommand.Cmd.AddCommand(
		profile.NewCreateProfileCommand(logger, profileCommand.Storage).Cmd,
		profile.NewDeleteProfileCommand(logger, profileCommand.Storage).Cmd,
		profile.NewGetProfileCommand(logger, profileCommand.Storage).Cmd,
		profile.NewListProfileCommand(logger, profileCommand.Storage).Cmd,
	)

	return profileCommand
}
