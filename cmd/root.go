package cmd

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rename-tool",
	Short: "A CLI tool for renaming and moving files",
	Long:  "Rename-tool automates renaming and moving files from a source directory to a target directory.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'rename-tool rename' to start renaming files.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		pterm.Error.Println(err)
		os.Exit(1)
	}
	pterm.Warning.Println("\nPress Enter to exit...")
	fmt.Scanln()
}
