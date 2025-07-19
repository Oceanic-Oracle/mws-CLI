package profile_storage

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"simple-cli/internal/config"
	"strings"

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

func (ps *profileStorage) List(path string) (string, error) {
	if path == "" {
		path = ps.cfg.File.Path
	}

	root := &treeNode{
		Path:      path,
		Name:      filepath.Base(path),
		IsDir:     true,
		ShowFiles: true,
	}

	if err := buildTree(root); err != nil {
		return "", fmt.Errorf("failed to build tree: %w", err)
	}

	return root.String(), nil
}

type treeNode struct {
	Path      string
	Name      string
	IsDir     bool
	Size      int64
	Children  []*treeNode
	ShowFiles bool
}

func buildTree(node *treeNode) error {
	entries, err := os.ReadDir(node.Path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		childPath := filepath.Join(node.Path, entry.Name())
		child := &treeNode{
			Path:      childPath,
			Name:      entry.Name(),
			IsDir:     entry.IsDir(),
			ShowFiles: node.ShowFiles,
		}

		if !entry.IsDir() {
			if info, err := entry.Info(); err == nil {
				child.Size = info.Size()
			}
		}

		node.Children = append(node.Children, child)
		if entry.IsDir() {
			if err := buildTree(child); err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *treeNode) String() string {
	var sb strings.Builder
	t.print(&sb, "", true)
	return sb.String()
}

func (t *treeNode) print(sb *strings.Builder, prefix string, isLast bool) {
	currentPrefix := prefix
	if isLast {
		currentPrefix += "└── "
	} else {
		currentPrefix += "├── "
	}

	sb.WriteString(currentPrefix + t.Name)

	if !t.IsDir {
		if t.Size == 0 {
			sb.WriteString(" (empty)")
		} else {
			sb.WriteString(fmt.Sprintf(" (%db)", t.Size))
		}
	}
	sb.WriteByte('\n')

	newPrefix := prefix
	if isLast {
		newPrefix += "    "
	} else {
		newPrefix += "│   "
	}

	for i, child := range t.Children {
		child.print(sb, newPrefix, i == len(t.Children)-1)
	}
}

func NewProfileStorage(cfg *config.Config, logger *slog.Logger) IProfile {
	profileStorage := &profileStorage{
		cfg:    cfg,
		logger: logger,
	}

	return profileStorage
}
