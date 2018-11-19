package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	// RootCmdName represents the cobra root command name.
	RootCmdName = "drm"

	// TODO replate it
	TempFileStorePath = "/home/lylex/workspace/drm/temp"
)

var (
	isForce     bool
	isRecursive bool
)

// RootCmd represents the base command.
var RootCmd = &cobra.Command{
	Use:   RootCmdName,
	Short: "A delayed rm with safety.",
	Long:  `This application is used to rm files with a latency.`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := os.Getwd()
		for _, item := range args {
			//	err := os.Rename(originalPath, newPath)
			currentPath := filepath.Join(dir, item)
			fmt.Println(currentPath)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Printf("Failed to execute command: %+v\n", err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.Flags().BoolVarP(&isRecursive, "recursive", "r", false, `ignore
nonexistent files and arguments, never prompt`)
	RootCmd.Flags().BoolVarP(&isForce, "force", "f", false, `remove directories
and their contents recursively or not`)
}
