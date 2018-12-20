package cmd

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/lylex/drm/pkg/files"

	"github.com/lylex/drm/pkg/blobs"
	"github.com/lylex/drm/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	gcCmdUse = "gc"

	// CfgTTLKey represents the ttl key in config, which means how many days after deletion we'd purge it.
	CfgTTLKey = "ttl"
)

var gcCmd = &cobra.Command{
	Use:   gcCmdUse,
	Short: "scan the metadata to find stale blobs and purge them",
	Long: `scan the metadata to find stale blobs and purge them, it should not be run by the user and
	it is created for the cron job`,
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		bs, err := blobs.GetAll()
		if err != nil {
			utils.ErrExit(fmt.Sprintf("\nFailed to retrieve data from database: %s\n", err.Error()), nil)
		}

		ttl, err := time.ParseDuration(viper.GetString(CfgTTLKey))
		if err != nil {
			utils.ErrExit("\nBad ttl configuration: %s\n", err)
		}
		expireDate := time.Now().Add(-ttl)

		for _, b := range bs {
			if b.CreatedAt.Before(expireDate) {
				if err := b.Destroy(); err != nil {
					utils.ErrExit(fmt.Sprintf("\nFailed to remove data from database: %s\n", err.Error()), nil)
				}
				files.Delete(filepath.Join(string(viper.GetString(CfgBlobPathKey)), b.Name()))
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(gcCmd)
}
