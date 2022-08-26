package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/DevelopNaoki/pcawui/api"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "pcawui server management commands",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("need valid command")
	},
}

var runCmd = &cobra.Command{
        Use:   "run",
        Short: "pcawui server run",
        RunE: func(cmd *cobra.Command, args []string) error {
		api.Run()
                return nil
        },
}
