package cmd

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/lylex/drm/pkg/files"

	"github.com/lylex/drm/pkg/blobs"
	"github.com/lylex/drm/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	restoreCmdName = "restore"
)

// restoreCmd represents the restore command.
var restoreCmd = &cobra.Command{
	Use:   restoreCmdName,
	Short: "restore one or more the deleted objects",
	Long:  `restore one or more the deleted objects, if a file has multiple version, the ID must be passed`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bs, err := blobs.GetAll()
		if err != nil {
			utils.ErrExit(fmt.Sprintf("\nFailed to retrieve data from database: %s\n", err.Error()), nil)
		}
		fileNameMap, idMap := index(bs)
		for _, arg := range args {
			if b, ok := idMap[arg]; ok {
				if err := handleBlob(b); err != nil {
					utils.ErrExit(fmt.Sprintf("\nFailed to restore object \"%s\"\n", arg), nil)
				}
				continue
			}

			if objs, ok := fileNameMap[arg]; ok {
				if len(objs) > 1 {
					utils.ErrExit(fmt.Sprintf("\nMore than one object named \"%s\", please use ID to restore it\n", arg), nil)
				}
				if err := handleBlob(objs[0]); err != nil {
					utils.ErrExit(fmt.Sprintf("\nFailed to restore object \"%s\"\n", arg), err)
				}
				continue
			}

			utils.ErrExit(fmt.Sprintf("\nNo file or directory with name or id \"%s\" found\n", arg), nil)
		}
	},
}

func handleBlob(b *blobs.Blob) error {
	if err := b.Destroy(); err != nil {
		return err
	}

	path := filepath.Join(b.Dir, b.FileName)
	if files.IsExist(path) {
		return errors.New("file or directory exists")
	}
	if err := files.Mkdir(b.Dir); err != nil {
		return err
	}
	files.Move(filepath.Join(string(viper.GetString(CfgBlobPathKey)), b.Name()), path)

	return nil
}

func index(bs []*blobs.Blob) (map[string][]*blobs.Blob, map[string]*blobs.Blob) {
	filenameMap := map[string][]*blobs.Blob{}
	idMap := map[string]*blobs.Blob{}

	for _, b := range bs {
		if _, ok := filenameMap[b.FileName]; ok {
			filenameMap[b.FileName] = append(filenameMap[b.FileName], b)
		} else {
			filenameMap[b.FileName] = []*blobs.Blob{b}
		}

		idMap[b.ID] = b
	}

	return filenameMap, idMap
}

func init() {
	RootCmd.AddCommand(restoreCmd)
}
