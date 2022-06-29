package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "pcawui",
	Short: "pcaui is private certificate authority web user interface",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("need valid command")
	},
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(
		serverCmd,
		testPkg,
	)
	serverCmd.AddCommand(
		serverStartCmd,
	)
}
