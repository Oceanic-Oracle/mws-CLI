package profile_test

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"simple-cli/cmd/commands/profile"
	"simple-cli/internal/config"
	"simple-cli/internal/logger"
	profile_storage "simple-cli/internal/storage/profile"
	"strings"
	"testing"
)

func TestCreateProfile(t *testing.T) {
	cfg := &config.Config{
        Log: struct{Level string `json:"level"`}{
			Level: "debug",
		},
		File: struct{Format string `json:"format"`; Path string `json:"path"`}{
			Format: "yaml",
			Path: t.TempDir(),
		},
    }
	logger := logger.SetupLogger(cfg.Log.Level)
	profileStorage := profile_storage.NewProfileStorage(cfg, logger)

	tests := []struct {
		name    string
		user    string
		project string
	}{
		{name: "profile1", user: "user1", project: "project1"},
		{name: "profile2", user: "user2", project: "project2"},
		{name: "duplicate", user: "user3", project: "project3"},
		{name: "test_profile4", user: "user4", project: "project4"},
		{name: "dev_profile5", user: "user5", project: "project5"},
		{name: "prod_profile6", user: "user6", project: "project6"},
		{name: "qa_profile7", user: "user7", project: "project7"},
		{name: "stage_profile8", user: "user8", project: "project8"},
		{name: "temp_profile9", user: "user9", project: "project9"},
		{name: "backup_profile10", user: "user10", project: "project10"},
		{name: "main_profile11", user: "user11", project: ""},
		{name: "secondary_profile12", user: "user12", project: "project12"},
		{name: "admin_profile13", user: "admin13", project: "project13"},
		{name: "guest_profile14", user: "guest14", project: "project14"},
		{name: "service_profile15", user: "service15", project: "project15"},
		{name: "api_profile16", user: "api16", project: "project16"},
		{name: "db_profile17", user: "dbadmin17", project: "project17"},
		{name: "log_profile18", user: "logger18", project: "project18"},
		{name: "batch_profile19", user: "", project: "project19"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            out := new(bytes.Buffer)
            
            profile.CreateProfile(out, profileStorage, "", tt.name, tt.user, tt.project)
            if out.String() != "" {
				t.Fatalf("%v", out.String())
			}

            expectedPath := filepath.Join(cfg.File.Path, fmt.Sprintf("%s.%s", tt.name, cfg.File.Format))
            if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
                t.Errorf("Profile file was not created: %s", expectedPath)
            }
            
            data, err := os.ReadFile(expectedPath)
            if err != nil {
                t.Fatalf("Failed to read created profile: %v", err)
            }
            
            if !strings.Contains(string(data), tt.user) {
                t.Errorf("Profile file doesn't contain user '%s'", tt.user)
            }
        })
	}
}

func TestDeleteProfile(t *testing.T) {
	cfg := &config.Config{
        Log: struct{Level string `json:"level"`}{
            Level: "debug",
        },
        File: struct{
            Format string `json:"format"`
            Path   string `json:"path"`
        }{
            Format: "yaml",
            Path:   t.TempDir(),
        },
    }
    logger := logger.SetupLogger(cfg.Log.Level)
    profileStorage := profile_storage.NewProfileStorage(cfg, logger)

    testProfiles := []struct {
        name    string
        user    string
        project string
    }{
        {name: "profile1", user: "user1", project: "project1"},
        {name: "profile2", user: "user2", project: "project2"},
        {name: "profile_to_delete", user: "user3", project: "project3"},
    }

    for _, tp := range testProfiles {
        out := new(bytes.Buffer)
        profile.CreateProfile(out, profileStorage, "", tp.name, tp.user, tp.project)
        if out.String() != "" {
            t.Fatalf("Failed to create test profile: %v", out.String())
        }
    }

    tests := []struct {
        name        string
        profileName string
        wantError   bool
        errorMsg    string
    }{
        {name: "delete existing profile", profileName: "profile_to_delete", wantError: false},
        {name: "delete non-existent profile", profileName: "nonexistent", wantError: true, errorMsg: fmt.Sprintf("failed to delete profile file: %s", filepath.Join(cfg.File.Path, fmt.Sprintf("%s.%s", "nonexistent", cfg.File.Format)))},
        {name: "delete already deleted profile", profileName: "profile_to_delete", wantError: true, errorMsg: fmt.Sprintf("failed to delete profile file: %s", filepath.Join(cfg.File.Path, fmt.Sprintf("%s.%s", "profile_to_delete", cfg.File.Format)))},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            out := new(bytes.Buffer)
            
            profile.DeleteProfile(out, profileStorage, "", tt.profileName)
            
            if tt.wantError {
                if out.String() == "" {
                    t.Error("Expected error message but got none")
                } else if tt.errorMsg != "" && !strings.Contains(out.String(), tt.errorMsg) {
                    t.Errorf("Expected error to contain %q, got %q", tt.errorMsg, out.String())
                }
            } else {
                if out.String() != "" {
                    t.Errorf("Unexpected error: %v", out.String())
                }
                
                expectedPath := filepath.Join(cfg.File.Path, fmt.Sprintf("%s.%s", tt.profileName, cfg.File.Format))
                if _, err := os.Stat(expectedPath); !os.IsNotExist(err) {
                    t.Errorf("Profile file still exists after deletion: %s", expectedPath)
                }
            }
        })
    }
}

func TestGetProfile(t *testing.T) {
    cfg := &config.Config{
        Log: struct{Level string `json:"level"`}{
            Level: "debug",
        },
        File: struct{
            Format string `json:"format"`
            Path   string `json:"path"`
        }{
            Format: "yaml",
            Path:   t.TempDir(),
        },
    }
    logger := logger.SetupLogger(cfg.Log.Level)
    profileStorage := profile_storage.NewProfileStorage(cfg, logger)

    testProfiles := []struct {
        name    string
        user    string
        project string
    }{
        {name: "profile1", user: "user1", project: "project1"},
        {name: "profile2", user: "user2", project: "project2"},
        {name: "empty_profile", user: "", project: ""},
    }

    for _, tp := range testProfiles {
        out := new(bytes.Buffer)
        profile.CreateProfile(out, profileStorage, "", tp.name, tp.user, tp.project)
        if out.String() != "" {
            t.Fatalf("Failed to create test profile: %v", out.String())
        }
    }

    tests := []struct {
        name        string
        profileName string
        wantOutput  string
    }{
        {name: "get existing profile", profileName: "profile1", wantOutput: "project: project1\nuser: user1\n"},
        {name: "get another profile", profileName: "profile2", wantOutput: "project: project2\nuser: user2\n"},
        {name: "get empty profile", profileName: "empty_profile", wantOutput: "project: \"\"\nuser: \"\"\n"},
        {name: "get non-existent profile", profileName: "nonexistent", wantOutput: fmt.Sprintf("failed to read profile file: open %s: The system cannot find the file specified.", filepath.Join(cfg.File.Path, fmt.Sprintf("%s.%s", "nonexistent", cfg.File.Format)))},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            out := new(bytes.Buffer)
            
            profile.GetProfile(out, profileStorage, "", tt.profileName)
            
            got := strings.ReplaceAll(out.String(), "\r\n", "\n")
            want := strings.ReplaceAll(tt.wantOutput, "\r\n", "\n")
            
            if got != want {
                t.Errorf("Expected output:\n%q\nbut got:\n%q", want, got)
            }
        })
    }
}


func TestListProfile(t *testing.T) {
    cfg := &config.Config{
        Log: struct{Level string `json:"level"`}{
            Level: "debug",
        },
        File: struct{
            Format string `json:"format"`
            Path   string `json:"path"`
        }{
            Format: "yaml",
            Path:   t.TempDir(),
        },
    }
    logger := logger.SetupLogger(cfg.Log.Level)
    profileStorage := profile_storage.NewProfileStorage(cfg, logger)

    testProfiles := []struct {
        name    string
        user    string
        project string
    }{
        {name: "dev", user: "dev1", project: "project1"},
        {name: "prod", user: "prod1", project: "project2"},
        {name: "test", user: "test1", project: "project3"},
    }

    for _, tp := range testProfiles {
        out := new(bytes.Buffer)
        profile.CreateProfile(out, profileStorage, "", tp.name, tp.user, tp.project)
        if out.String() != "" {
            t.Fatalf("Failed to create test profile: %v", out.String())
        }
    }

    nonProfileFiles := []struct {
        name    string
        content string
    }{
        {name: "readme.md", content: "test"},
        {name: "config.json", content: `{"key":"value"}`},
        {name: "temp.txt", content: "temp"},
    }

    for _, f := range nonProfileFiles {
        filePath := filepath.Join(cfg.File.Path, f.name)
        if err := os.WriteFile(filePath, []byte(f.content), 0644); err != nil {
            t.Fatalf("Failed to create test file: %v", err)
        }
    }

    tests := []struct {
        name       string
        path       string
        wantOutput string
    }{
        {
            name: "list profiles",
            path: "",
            wantOutput: "└── " + filepath.Base(cfg.File.Path) + "\n" +
                "    ├── config.json (15b)\n" +
                "    ├── dev.yaml (29b)\n" +
                "    ├── prod.yaml (30b)\n" +
                "    ├── readme.md (4b)\n" +
                "    ├── temp.txt (4b)\n" +
                "    └── test.yaml (30b)\n",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            out := new(bytes.Buffer)
            
            profile.ListProfile(out, profileStorage, tt.path)
            
            got := strings.ReplaceAll(out.String(), "\r\n", "\n")
            want := strings.ReplaceAll(tt.wantOutput, "\r\n", "\n")
            
            if got != want {
                t.Errorf("Expected output:\n%q\nbut got:\n%q", want, got)
                t.Logf("Full output:\n%s", out.String())
            }
        })
    }
}
