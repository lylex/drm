package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/lylex/drm/pkg/files"
	"github.com/lylex/drm/pkg/utils"
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
// Actually, we do not valid the sub-command here since we need to execute something
// like:
//     drm test.txt
//     drm -f test.txt
//     drm ls
// So we have to attept ArbitraryArgs.
var RootCmd = &cobra.Command{
	Use:   RootCmdName,
	Short: "A delayed rm with safety.",
	Long:  `This application is used to rm files with a latency.`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		for _, item := range args {
			dir := files.GetWd()
			currentPath := filepath.Join(dir, item)
			fmt.Println(currentPath)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		utils.PrintErr("Failed to execute command: %+v\n", err)
	}
}

func init() {
	RootCmd.Flags().BoolVarP(&isRecursive, "recursive", "r", false, `ignore
nonexistent files and arguments, never prompt`)
	RootCmd.Flags().BoolVarP(&isForce, "force", "f", false, `remove directories
and their contents recursively or not`)
}
