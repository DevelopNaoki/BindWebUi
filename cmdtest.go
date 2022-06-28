package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/DevelopNaoki/pcawui/api"
)

var testPkg = &cobra.Command{
	Use:   "test",
	Short: "Package test",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("need valid command")
	},
}
