package cmd

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/lylex/drm/pkg/files"
	"github.com/lylex/drm/pkg/utils"
	"github.com/spf13/cobra"
)

const (
	// RootCmdName represents the cobra root command name.
	RootCmdName = "drm"

	// TODO replate it
	TempFileStorePath = "/Users/xuq3/workspace/drm/tem"
)

var (
	isForce     bool
	isRecursive bool
	version     string = "1.2"
)

func init() {
	RootCmd.Flags().BoolVarP(&isRecursive, "recursive", "r", false, `ignore
nonexistent files and arguments, never prompt`)
	RootCmd.PersistentFlags().BoolVarP(&isForce, "force", "f", false, `remove directories
and their contents recursively or not`)
}

// RootCmd represents the base command.
// Actually, we do not valid the sub-command here since we need to execute something
// like:
//     drm test.txt
//     drm -f test.txt
//     drm ls
// So we have to attept ArbitraryArgs.
var RootCmd = &cobra.Command{
	Use:     RootCmdName,
	Short:   "A delayed rm with safety.",
	Long:    `This application is used to rm files with a latency.`,
	Version: fmt.Sprintf("%s", version),
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		for _, path := range args {
			if !files.IsAbsolutePath(path) {
				path = filepath.Join(files.GetWd(), path)
			}

			if !files.IsExist(path) {
				if isForce {
					continue
				}
				utils.ErrExit(fmt.Sprintf("%s: %s: No such file or directory\n", RootCmdName, path), nil)
			}

			if files.IsDir(path) && !isRecursive {
				utils.ErrExit(fmt.Sprintf("%s: %s: is a directory\n", RootCmdName, path), nil)
			}
			files.Move(path, filepath.Join(TempFileStorePath,
				fmt.Sprintf("%d_%s", time.Now().UnixNano(), files.Name(path))))
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		utils.ErrExit("Failed to execute command: %+v\n", err)
	}
}
