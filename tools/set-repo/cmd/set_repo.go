// Package cmd
package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gtag"
)

// SetRepo is the management object for the `scaffold new` sub-command.
var SetRepo = cSetRepo{}

type cSetRepo struct {
	g.Meta `name:"set-repo" brief:"{cSetRepoBrief}" eg:"{cSetRepoEg}"`
}

const (
	cSetRepoBrief = `change the import paths to match the new module name`
	cSetRepoEg    = `set-repo repo-name
set-repo github.com/user/repo
`
)

const (
	cSetRepoModuleBrief = `name of the new module (kebab-case prefered, e.g. reso, my-repo, example-service)`
)

func init() {
	gtag.Sets(g.MapStrStr{
		`cSetRepoBrief`:       cSetRepoBrief,
		`cSetRepoEg`:          cSetRepoEg,
		`cSetRepoModuleBrief`: cSetRepoModuleBrief,
	})
}

type cNewInput struct {
	g.Meta `       name:"set-repo"`
	Name   string `name:"NAME"     arg:"true" brief:"{cSetRepoModuleBrief}"`
}

type cNewOutput struct{}

// Index implements the `scaffold new` command.
func (c cSetRepo) Index(ctx context.Context, in cNewInput) (out *cNewOutput, err error) {
	g.Dump(in)
	// oldName, _ := pterm.DefaultInteractiveTextInput.Show("enter old name")
	//
	// pterm.Println()
	// pterm.Info.Printfln("You answered: %s", oldName)

	// newName, _ := pterm.DefaultInteractiveTextInput.Show("enter new name")
	runCmd(`qistas/`, `modular-monolith/`)
	// if err != nil {
	// 	return
	// }

	return
}

func runCmd(old, new string) {
	root := "./tools_1/scaffold"
	// old = `qistas/`
	// new = `modular-monolith/`
	if err := replaceInDirectory(root, old, new); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func replaceInDirectory(root, oldText, newText string) error {
	return filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories.
		if d.IsDir() {
			return nil
		}

		// Only process regular files.
		info, err := d.Info()
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read %s: %w", path, err)
		}

		// Skip files that don't contain the text.
		if !strings.Contains(string(content), oldText) {
			return nil
		}

		updated := strings.ReplaceAll(string(content), oldText, newText)

		// Preserve the original permissions.
		if err := os.WriteFile(path, []byte(updated), info.Mode().Perm()); err != nil {
			return fmt.Errorf("write %s: %w", path, err)
		}

		fmt.Println("Updated:", path)

		return nil
	})
}
