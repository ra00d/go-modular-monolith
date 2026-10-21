package db

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

const testMigrationTemplate = `-- +goose Up
SELECT 'up SQL query';

-- +goose Down
SELECT 'down SQL query';
`

func TestCreateNewMigration(t *testing.T) {
	tests := []struct {
		name          string
		migrationName string
		path          string
		wantSuffix    string
		wantErr       error
	}{
		{
			name:          "creates migration in relative path",
			migrationName: "Add User Preferences",
			path:          "internal/auth/migrations",
			wantSuffix:    "add_user_preferences.sql",
		},
		{
			name:          "removes sql extension from name",
			migrationName: "add_users.sql",
			path:          "migrations",
			wantSuffix:    "add_users.sql",
		},
		{
			name:          "rejects empty name",
			migrationName: "   ",
			path:          "migrations",
			wantErr:       ErrInvalidMigrationName,
		},
		{
			name:          "rejects name without letters or digits",
			migrationName: "---",
			path:          "migrations",
			wantErr:       ErrInvalidMigrationName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := t.TempDir()
			writeTemplate(t, root)
			t.Chdir(root)

			err := CreateNewMigration(tt.migrationName, tt.path)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("CreateNewMigration() error = %v, want %v", err, tt.wantErr)
				}

				return
			}

			if err != nil {
				t.Fatalf("CreateNewMigration() error = %v", err)
			}

			matches, err := filepath.Glob(filepath.Join(tt.path, "*_"+tt.wantSuffix))
			if err != nil {
				t.Fatalf("glob migration file: %v", err)
			}
			if len(matches) != 1 {
				t.Fatalf("found %d migration files matching %q, want 1", len(matches), "*_"+tt.wantSuffix)
			}

			content, err := os.ReadFile(matches[0])
			if err != nil {
				t.Fatalf("read migration file: %v", err)
			}
			if string(content) != testMigrationTemplate {
				t.Fatalf("migration content = %q, want %q", content, testMigrationTemplate)
			}
		})
	}
}

func TestCreateNewMigrationMissingTemplate(t *testing.T) {
	root := t.TempDir()
	t.Chdir(root)

	err := CreateNewMigration("add_users", "migrations")
	if err == nil {
		t.Fatal("CreateNewMigration() error = nil, want an error for a missing template")
	}
}

func writeTemplate(t *testing.T, root string) {
	t.Helper()

	templatePath := filepath.Join(root, migrationTemplatePath)
	if err := os.MkdirAll(filepath.Dir(templatePath), 0o755); err != nil {
		t.Fatalf("create template directory: %v", err)
	}
	if err := os.WriteFile(templatePath, []byte(testMigrationTemplate), 0o644); err != nil {
		t.Fatalf("write template: %v", err)
	}
}
