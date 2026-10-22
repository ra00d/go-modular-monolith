// Package modulegen generates new application modules from the templates in
// resource/templates.
package modulegen

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"time"
	"unicode"

	"modular-monolith/lib/utils"
)

const (
	templatesDir = "resource/templates"

	moduleTemplate       = "module.go.tmpl"
	appTemplate          = "app.go.tmpl"
	restTemplate         = "rest.go.tmpl"
	migrationsTemplate   = "migrations.go.tmpl"
	migrationSQLTemplate = "migration.sql.tmpl"
)

var (
	// ErrInvalidModuleName is returned when a module name cannot be used as a
	// Go package name.
	ErrInvalidModuleName = errors.New(
		"module name must start with a letter and contain only letters, digits, and underscores",
	)
	// ErrModuleExists is returned when the destination module directory already
	// exists.
	ErrModuleExists = errors.New("module directory already exists")
)

var moduleNamePattern = regexp.MustCompile(`^[a-z][a-z0-9_]*$`)

// Options controls where a module is generated.
type Options struct {
	// Dir is the module destination, relative to the current working directory
	// or absolute. When empty, internal/modules/<name> is used if that directory
	// exists, otherwise internal/<name>.
	Dir string
}

type moduleData struct {
	ModuleName  string
	ModuleTitle string
	GoModule    string
	ModulePath  string
}

type generatedFile struct {
	templateName string
	path         string
}

// CreateModule scaffolds a module with its REST, application, domain, infra,
// and migrations packages.
func CreateModule(name string, opts Options) error {
	moduleName, err := NormalizeModuleName(name)
	if err != nil {
		return err
	}

	workingDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get current working directory: %w", err)
	}

	moduleDir, err := resolveModuleDir(workingDir, moduleName, opts.Dir)
	if err != nil {
		return err
	}

	if _, err := os.Stat(moduleDir); err == nil {
		return fmt.Errorf("%w: %s", ErrModuleExists, moduleDir)
	} else if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("check module directory %q: %w", moduleDir, err)
	}

	goModule, err := utils.ReadGoModulePath(workingDir)
	if err != nil {
		return err
	}

	modulePath, err := filepath.Rel(workingDir, moduleDir)
	if err != nil {
		return fmt.Errorf("resolve module path: %w", err)
	}
	modulePath = filepath.ToSlash(modulePath)

	data := moduleData{
		ModuleName:  moduleName,
		ModuleTitle: title(moduleName),
		GoModule:    goModule,
		ModulePath:  modulePath,
	}

	dirs := []string{
		filepath.Join(moduleDir, "api", "rest"),
		filepath.Join(moduleDir, "app"),
		filepath.Join(moduleDir, "domain"),
		filepath.Join(moduleDir, "infra"),
		filepath.Join(moduleDir, "migrations"),
	}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("create directory %q: %w", dir, err)
		}
	}

	files := []generatedFile{
		{moduleTemplate, filepath.Join(moduleDir, "module.go")},
		{appTemplate, filepath.Join(moduleDir, "app", "app.go")},
		{restTemplate, filepath.Join(moduleDir, "api", "rest", "rest.go")},
		{migrationsTemplate, filepath.Join(moduleDir, "migrations", "migrations.go")},
	}
	for _, file := range files {
		if err := renderTemplate(workingDir, file.templateName, file.path, data); err != nil {
			return err
		}
	}

	// migrations.go embeds *.sql, so the package needs at least one migration
	// to compile.
	if err := createInitialMigration(workingDir, filepath.Join(moduleDir, "migrations")); err != nil {
		return err
	}

	// Keep intentionally empty package directories in version control.
	for _, dir := range []string{"domain", "infra"} {
		path := filepath.Join(moduleDir, dir, ".gitkeep")
		if err := os.WriteFile(path, nil, 0o644); err != nil {
			return fmt.Errorf("create placeholder %q: %w", path, err)
		}
	}

	return nil
}

// NormalizeModuleName lowercases a module name and replaces spaces and hyphens
// with underscores.
func NormalizeModuleName(name string) (string, error) {
	var b strings.Builder
	lastUnderscore := false

	for _, r := range strings.ToLower(strings.TrimSpace(name)) {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			b.WriteRune(r)
			lastUnderscore = false
		case unicode.IsSpace(r) || r == '-' || r == '_':
			if !lastUnderscore && b.Len() > 0 {
				b.WriteRune('_')
				lastUnderscore = true
			}
		default:
			return "", ErrInvalidModuleName
		}
	}

	moduleName := strings.Trim(b.String(), "_")
	if !moduleNamePattern.MatchString(moduleName) {
		return "", ErrInvalidModuleName
	}

	return moduleName, nil
}

func resolveModuleDir(workingDir, moduleName, configuredDir string) (string, error) {
	var moduleDir string
	if configuredDir != "" {
		moduleDir = configuredDir
		if !filepath.IsAbs(moduleDir) {
			moduleDir = filepath.Join(workingDir, moduleDir)
		}
	} else {
		modulesDir := filepath.Join(workingDir, "internal", "modules")
		if info, err := os.Stat(modulesDir); err == nil && info.IsDir() {
			moduleDir = filepath.Join(modulesDir, moduleName)
		} else {
			moduleDir = filepath.Join(workingDir, "internal", moduleName)
		}
	}

	moduleDir = filepath.Clean(moduleDir)

	rel, err := filepath.Rel(workingDir, moduleDir)
	if err != nil {
		return "", fmt.Errorf("resolve module directory: %w", err)
	}
	if rel == "." || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("module directory must be inside the project root: %s", configuredDir)
	}

	return moduleDir, nil
}

func renderTemplate(workingDir, templateName, destination string, data moduleData) error {
	templatePath := filepath.Join(workingDir, templatesDir, templateName)
	content, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("read template %q: %w", templatePath, err)
	}

	tmpl, err := template.New(templateName).Option("missingkey=error").Parse(string(content))
	if err != nil {
		return fmt.Errorf("parse template %q: %w", templatePath, err)
	}

	var rendered bytes.Buffer
	if err := tmpl.Execute(&rendered, data); err != nil {
		return fmt.Errorf("render template %q: %w", templatePath, err)
	}

	if err := os.WriteFile(destination, rendered.Bytes(), 0o644); err != nil {
		return fmt.Errorf("write module file %q: %w", destination, err)
	}

	return nil
}

func createInitialMigration(workingDir, migrationDir string) error {
	templatePath := filepath.Join(workingDir, templatesDir, migrationSQLTemplate)
	content, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("read template %q: %w", templatePath, err)
	}

	fileName := fmt.Sprintf("%s_init.sql", time.Now().Format("20060102150405"))
	migrationPath := filepath.Join(migrationDir, fileName)
	if err := os.WriteFile(migrationPath, content, 0o644); err != nil {
		return fmt.Errorf("write migration file %q: %w", migrationPath, err)
	}

	return nil
}

func title(value string) string {
	runes := []rune(value)
	if len(runes) == 0 {
		return value
	}
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
