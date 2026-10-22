package modulegen

import (
	"errors"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCreateModule(t *testing.T) {
	root := setupProject(t)

	if err := CreateModule("Sales Orders", Options{}); err != nil {
		t.Fatalf("CreateModule() error = %v", err)
	}

	moduleDir := filepath.Join(root, "internal", "sales_orders")
	for _, dir := range []string{
		filepath.Join(moduleDir, "api", "rest"),
		filepath.Join(moduleDir, "app"),
		filepath.Join(moduleDir, "domain"),
		filepath.Join(moduleDir, "infra"),
		filepath.Join(moduleDir, "migrations"),
	} {
		info, err := os.Stat(dir)
		if err != nil {
			t.Fatalf("stat %q: %v", dir, err)
		}
		if !info.IsDir() {
			t.Fatalf("%q is not a directory", dir)
		}
	}

	files := []string{
		filepath.Join(moduleDir, "module.go"),
		filepath.Join(moduleDir, "app", "app.go"),
		filepath.Join(moduleDir, "api", "rest", "rest.go"),
		filepath.Join(moduleDir, "migrations", "migrations.go"),
	}
	for _, file := range files {
		assertGoFile(t, file)
	}

	moduleGo := readFile(t, filepath.Join(moduleDir, "module.go"))
	assertContains(t, moduleGo, "package sales_orders")
	assertContains(t, moduleGo, `"example.com/app/internal/sales_orders/api/rest"`)
	assertContains(t, moduleGo, `"example.com/app/internal/sales_orders/migrations"`)
	assertContains(t, moduleGo, `modules.Register(Module, "Sales_orders")`)

	restGo := readFile(t, filepath.Join(moduleDir, "api", "rest", "rest.go"))
	assertContains(t, restGo, `srv.Group("/sales_orders"`)

	migrations, err := filepath.Glob(filepath.Join(moduleDir, "migrations", "*.sql"))
	if err != nil {
		t.Fatalf("glob migrations: %v", err)
	}
	if len(migrations) != 1 {
		t.Fatalf("found %d migration files, want 1", len(migrations))
	}
	assertGoFile(t, migrations[0])
}

func TestCreateModuleUsesModulesDirectoryWhenPresent(t *testing.T) {
	root := setupProject(t)
	if err := os.MkdirAll(filepath.Join(root, "internal", "modules"), 0o755); err != nil {
		t.Fatalf("create modules directory: %v", err)
	}

	if err := CreateModule("orders", Options{}); err != nil {
		t.Fatalf("CreateModule() error = %v", err)
	}

	if _, err := os.Stat(filepath.Join(root, "internal", "modules", "orders", "module.go")); err != nil {
		t.Fatalf("expected module under internal/modules: %v", err)
	}
}

func TestCreateModuleWithCustomPath(t *testing.T) {
	root := setupProject(t)

	if err := CreateModule("orders", Options{Dir: "internal/custom/orders"}); err != nil {
		t.Fatalf("CreateModule() error = %v", err)
	}

	moduleGo := readFile(t, filepath.Join(root, "internal", "custom", "orders", "module.go"))
	assertContains(t, moduleGo, `"example.com/app/internal/custom/orders/migrations"`)
}

func TestCreateModuleRefusesToOverwrite(t *testing.T) {
	root := setupProject(t)
	if err := os.MkdirAll(filepath.Join(root, "internal", "orders"), 0o755); err != nil {
		t.Fatalf("create module directory: %v", err)
	}

	err := CreateModule("orders", Options{})
	if !errors.Is(err, ErrModuleExists) {
		t.Fatalf("CreateModule() error = %v, want %v", err, ErrModuleExists)
	}
}

func TestCreateModuleRejectsOutsideProject(t *testing.T) {
	setupProject(t)

	err := CreateModule("orders", Options{Dir: t.TempDir()})
	if err == nil {
		t.Fatal("CreateModule() error = nil, want an error for a directory outside the project")
	}
}

func TestNormalizeModuleName(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		{name: "lowercase name", input: "orders", want: "orders"},
		{name: "spaces and case", input: "Sales Orders", want: "sales_orders"},
		{name: "hyphen", input: "sales-orders", want: "sales_orders"},
		{name: "repeated separators", input: " sales--orders ", want: "sales_orders"},
		{name: "empty", input: "   ", wantErr: ErrInvalidModuleName},
		{name: "starts with digit", input: "1orders", wantErr: ErrInvalidModuleName},
		{name: "invalid character", input: "sales/orders", wantErr: ErrInvalidModuleName},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalizeModuleName(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("NormalizeModuleName() error = %v, want %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("NormalizeModuleName() error = %v", err)
			}
			if got != tt.want {
				t.Fatalf("NormalizeModuleName() = %q, want %q", got, tt.want)
			}
		})
	}
}

func setupProject(t *testing.T) string {
	t.Helper()

	templatesPath, err := filepath.Abs(filepath.Join("..", "..", "..", "resource", "templates"))
	if err != nil {
		t.Fatalf("resolve templates path: %v", err)
	}

	root := t.TempDir()
	if err := os.CopyFS(filepath.Join(root, "resource", "templates"), os.DirFS(templatesPath)); err != nil {
		t.Fatalf("copy templates: %v", err)
	}
	if err := os.WriteFile(
		filepath.Join(root, "go.mod"),
		[]byte("module example.com/app\n\ngo 1.25.0\n"),
		0o644,
	); err != nil {
		t.Fatalf("write go.mod: %v", err)
	}

	t.Chdir(root)

	return root
}

func assertGoFile(t *testing.T, path string) {
	t.Helper()

	if !strings.HasSuffix(path, ".go") {
		return
	}
	if _, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.AllErrors); err != nil {
		t.Fatalf("parse %q: %v", path, err)
	}
}

func readFile(t *testing.T, path string) string {
	t.Helper()

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %q: %v", path, err)
	}

	return string(content)
}

func assertContains(t *testing.T, value, want string) {
	t.Helper()

	if !strings.Contains(value, want) {
		t.Fatalf("value does not contain %q:\n%s", want, value)
	}
}
