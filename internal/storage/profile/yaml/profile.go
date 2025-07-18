package profile_yaml

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"simple-cli/internal/config"
	profile_storage "simple-cli/internal/storage/profile"

	"sigs.k8s.io/yaml"
)

type profileStorage struct {
	cfg    *config.Config
	logger *slog.Logger
}

func (ps *profileStorage) Create(outPath, name, user, project string) error {
	if outPath == "" {
		outPath = ps.cfg.File.Path
	}
	
	file, err := os.Create(path.Join(outPath, fmt.Sprint(name, ".", ps.cfg.File.Format)))
	if err != nil {
		return err
	}
	defer file.Close()

	profile := profile_storage.Profile{
		User:    user,
		Project: project,
	}

	data, err := yaml.Marshal(profile)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func NewProfileStorage(cfg *config.Config, logger *slog.Logger) profile_storage.IProfile {
	profileStorage := &profileStorage{
		cfg:    cfg,
		logger: logger,
	}

	return profileStorage
}
