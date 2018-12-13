package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	lsCmdUse = "ls"
)

// lsCmd represents the ls command.
var lsCmd = &cobra.Command{
	Use:   lsCmdUse,
	Short: "list all the deleted objects",
	Long:  `list all the deleted objects, and all can be resumed`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("in the ls cmd")
	},
}

func init() {
	RootCmd.AddCommand(lsCmd)
}
