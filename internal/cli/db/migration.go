package db

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"

	"modular-monolith/lib/config"
	libdb "modular-monolith/lib/db"
)

const migrationTemplatePath = "resource/templates/migration.sql.tmpl"

var ErrInvalidMigrationName = errors.New(
	"migration name must contain at least one letter or digit",
)

// CreateNewMigration creates a new Goose migration from the SQL template.
// Both the destination path and the template location are resolved relative
// to the current working directory.
func CreateNewMigration(name, module string) error {
	fileName, err := migrationFileName(name)
	if err != nil {
		return err
	}

	workingDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get current working directory: %w", err)
	}

	templatePath := filepath.Join(workingDir, migrationTemplatePath)
	template, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("read migration template %q: %w", templatePath, err)
	}

	migrationDir := "internal/" + module + "/migrations"
	if !filepath.IsAbs(migrationDir) {
		migrationDir = filepath.Join(workingDir, migrationDir)
	}

	if err := os.MkdirAll(migrationDir, 0o755); err != nil {
		return fmt.Errorf("create migration directory %q: %w", migrationDir, err)
	}

	migrationPath := filepath.Join(migrationDir, fileName)

	file, err := os.OpenFile(migrationPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		return fmt.Errorf("create migration file %q: %w", migrationPath, err)
	}

	if _, err := file.Write(template); err != nil {
		_ = file.Close()
		return fmt.Errorf("write migration file %q: %w", migrationPath, err)
	}

	if err := file.Close(); err != nil {
		return fmt.Errorf("close migration file %q: %w", migrationPath, err)
	}

	return nil
}

func MigrateDB(module string) error {
	genPath := "internal/" + module
	dsn, err := config.LoadConfig().Database.DSNWithSchema(module)
	if err != nil {
		return err
	}
	migrations := os.DirFS("internal/" + module + "/migrations")
	libdb.MigrateDB(module, genPath, dsn, migrations)
	return nil
}

// migrationFileName returns a timestamped, filesystem-safe migration name.
func migrationFileName(name string) (string, error) {
	name = strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(name), ".sql"))

	var b strings.Builder
	lastUnderscore := false

	for _, r := range strings.ToLower(name) {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			b.WriteRune(r)
			lastUnderscore = false
		case unicode.IsSpace(r) || r == '-' || r == '_':
			if !lastUnderscore && b.Len() > 0 {
				b.WriteRune('_')
				lastUnderscore = true
			}
		}
	}

	cleanName := strings.Trim(b.String(), "_")
	if cleanName == "" {
		return "", ErrInvalidMigrationName
	}

	return fmt.Sprintf("%s_%s.sql", time.Now().Format("20060102150405"), cleanName), nil
}
