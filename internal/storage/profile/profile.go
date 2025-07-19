package profile_storage

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"simple-cli/internal/config"

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

	filePath := path.Join(outPath, fmt.Sprintf("%s.%s", name, ps.cfg.File.Format))

	if _, err := os.Stat(filePath); err == nil {
		return fmt.Errorf("profile file '%s' already exists", filePath)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("failed to check profile existence: %w", err)
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return fmt.Errorf("failed to create profile file: %w", err)
	}
	defer file.Close()

	profile := Profile{
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

func (ps *profileStorage) Delete(outPath, name string) error {
	if outPath == "" {
		outPath = ps.cfg.File.Path
	}

	if err := os.Remove(path.Join(outPath, fmt.Sprintf("%s.%s", name, ps.cfg.File.Format))); err != nil {
		return fmt.Errorf("failed to delete profile file: %w", err)
	}

	return nil
}

func (ps *profileStorage) Get(outPath, name string) (string, error) {
	if outPath == "" {
		outPath = ps.cfg.File.Path
	}

	data, err := os.ReadFile(path.Join(outPath, fmt.Sprintf("%s.%s", name, ps.cfg.File.Format)))
	if err != nil {
		return "", fmt.Errorf("failed to read profile file: %w", err)
	}

	return string(data), nil
}

func NewProfileStorage(cfg *config.Config, logger *slog.Logger) IProfile {
	profileStorage := &profileStorage{
		cfg:    cfg,
		logger: logger,
	}

	return profileStorage
}
