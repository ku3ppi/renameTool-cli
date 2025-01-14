// CLi version
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var sourceDir, targetDir, baseName string

var renameCmd = &cobra.Command{
	Use:     "rename",
	Aliases: []string{"r", "rn"},
	Short:   "Rename and move files",
	Run: func(cmd *cobra.Command, args []string) {
		if sourceDir == "" || targetDir == "" || baseName == "" {
			pterm.Error.Println("sourceDir, targetDir, and baseName are required")
			return
		}

		renameAndMoveFiles(sourceDir, targetDir, baseName)
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
	renameCmd.Flags().StringVarP(&sourceDir, "source", "s", "", "Source directory (required)")
	renameCmd.Flags().StringVarP(&targetDir, "target", "t", "", "Target directory (required)")
	renameCmd.Flags().StringVarP(&baseName, "name", "n", "", "Base name for renamed files (required)")
}

func renameAndMoveFiles(sourceDir, targetDir, baseName string) {
	supportedFormats := []string{".jpg", ".jpeg", ".CR2", ".CR3", ".CRW"}
	files, err := os.ReadDir(sourceDir)
	if err != nil {
		pterm.Error.Printf("Error reading source directory: %v\n", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file.Name()))
		isSupported := false
		for _, format := range supportedFormats {
			if ext == format {
				isSupported = true
				break
			}
		}
		if !isSupported {
			continue
		}

		fileID := strings.TrimSuffix(file.Name(), ext)
		newFilename := fmt.Sprintf("%s_%s%s", baseName, fileID, ext)
		oldPath := filepath.Join(sourceDir, file.Name())
		newPath := filepath.Join(targetDir, newFilename)

		err := os.Rename(oldPath, newPath)
		if err != nil {
			pterm.Warning.Printf("Failed to move: %s -> %s\n", oldPath, newPath)
			continue
		}
		pterm.Success.Printf("Moved: %s -> %s\n", oldPath, newPath)
	}
}
