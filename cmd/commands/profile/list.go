package profile

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	profile_storage "simple-cli/internal/storage/profile"

	"github.com/spf13/cobra"
)

type ListProfileCommand struct {
	logger *slog.Logger
	Cmd    *cobra.Command
}

func NewListProfileCommand(logger *slog.Logger, profileStorage profile_storage.IProfile) *ListProfileCommand {
	listProfileCommand := &ListProfileCommand{
		logger: logger,
		Cmd: &cobra.Command{
			Use:   "list",
			Short: "List all available profiles",
			Run: func(cmd *cobra.Command, args []string) {
				var path string
				if len(args) > 0 {
					path = args[0]
				}

				ListProfile(os.Stdout, profileStorage, path)
			},
		},
	}

	return listProfileCommand
}

// ListProfile выводит список профилей (экспортируемая для тестов)
func ListProfile(out io.Writer, profileStorage profile_storage.IProfile, path string) {
	data, err := profileStorage.List(path)
	if err != nil {
		fmt.Fprintf(out, "%v", err)
		return
	}

	fmt.Fprintf(out, "%s", data)
}
