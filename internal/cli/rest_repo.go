package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"modular-monolith/lib/utils"

	"github.com/spf13/cobra"
)

var (
	rootDir      = "./test-data"
	resetRepoCmd = &cobra.Command{
		Use:   "set-name [name]",
		Short: "change the module name",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// rootDir = utils.Deref(
			// 	cmd.PersistentFlags().
			// 		String("rootDir", ".", "The root directory of the project"),
			// )
			oldName := utils.GetGoModulePath()
			resetRepo(oldName, args[0])
		},
	}
)

func init() {
	resetRepoCmd.PersistentFlags().
		StringVarP(&rootDir, "root", "r", ".", "The root directory of the project")

	rootCmd.AddCommand(resetRepoCmd)
}

func resetRepo(old, new string) {
	root := rootDir
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

func firstNWords(s string, n int) string {
	words := strings.Fields(s)
	if len(words) <= n {
		return strings.Join(words, " ")
	}
	return strings.Join(words[:n], " ")
}
