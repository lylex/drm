package cmd

import (
	"fmt"

	"github.com/cheynewallace/tabby"
	"github.com/lylex/drm/pkg/blobs"
	"github.com/lylex/drm/pkg/utils"
	"github.com/spf13/cobra"
)

const (
	listCmdName = "list"

	timeFormat = "2006-01-02 15:04:05"
)

type table struct {
	*tabby.Tabby
}

func createTable() *table {
	return &table{
		tabby.New(),
	}
}

func (t *table) addHeaders(headers ...interface{}) {
	t.AddHeader(headers...)
}

func (t *table) addRow(items ...interface{}) {
	t.AddLine(items...)
}

// listCmd represents the list command.
var listCmd = &cobra.Command{
	Use:   listCmdName,
	Short: "list all the deleted objects",
	Long:  `list all the deleted objects, and all can be restored, try "drm restore" to restore an object`,
	Run: func(cmd *cobra.Command, args []string) {
		table := createTable()
		table.addHeaders("Name", "Path", "DeleteAt", "ID")

		blobs, err := blobs.GetAll()
		if err != nil {
			utils.ErrExit(fmt.Sprintf("%s %s: failed to retrieve data from database\n", RootCmdName, listCmdName), err)
		}
		for _, blob := range blobs {
			table.addRow(blob.FileName, blob.Dir, blob.CreatedAt.Format(timeFormat), blob.ID)
		}
		table.Print()
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
