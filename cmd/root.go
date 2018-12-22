package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lylex/drm/pkg/blobs"
	"github.com/lylex/drm/pkg/files"
	"github.com/lylex/drm/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// RootCmdName represents the cobra root command name.
	RootCmdName = "drm"

	// DefaultConfigPath represents the default path for config.
	DefaultConfigPath = "/etc/drm/drm.conf"

	// CfgBlobPathKey represents the blobPath key in config.
	CfgBlobPathKey = "blobPath"
)

var (
	isForce     bool
	isRecursive bool
	version     string
	cfgFile     = DefaultConfigPath
)

func init() {
	cobra.OnInitialize(initConfig)

	// determin whether to remove recursively.
	RootCmd.Flags().BoolVarP(&isRecursive, "recursive", "r", false, "remove directories and their contents recursively or not")

	// determine whether to remove objects by force.
	RootCmd.Flags().BoolVarP(&isForce, "force", "f", false, "ignore nonexistent files and arguments, never prompt")

	// pass configuration file from options.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", DefaultConfigPath,
		fmt.Sprintf("config file (default is %s)", DefaultConfigPath))
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
		if len(args) == 0 {
			msg := "no operation concluded"
			fmt.Fprintf(os.Stderr, fmt.Sprintf("Error: %s\n", msg))
			cmd.Usage()
			utils.ErrExit(fmt.Sprintf("\nFailed to execute command: %s\n", msg), nil)
		}

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

			blob := blobs.Create(path)
			files.Move(path, filepath.Join(string(viper.GetString(CfgBlobPathKey)), blob.Name()))
			if err := blob.Save(); err != nil {
				utils.ErrExit(fmt.Sprintf("%s: failed to save metadata for %s", RootCmdName, path), nil)
			}
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

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile == "" {
		cfgFile = DefaultConfigPath
	}
	viper.SetConfigFile(cfgFile)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		utils.ErrExit("Failed to read configuration file: %+v\n", err)
	}

	// read config from environment virable.
	viper.AutomaticEnv()
}
