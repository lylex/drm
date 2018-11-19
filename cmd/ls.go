package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	lsCmdUse = "ls"
)

// lsCmd represents the ls command.
// TODO fill the TBDs
var lsCmd = &cobra.Command{
	Use:   lsCmdUse,
	Short: "short TBD",
	Long:  `long TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("in the ls cmd")
	},
}

func init() {
	RootCmd.AddCommand(lsCmd)
}
