package cmd

import (
	"github.com/spf13/cobra"
)

// startCMD represents the start command
var startCMD = &cobra.Command{
	Use:   "start",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return startHttpGateway()
	},
}

func init() {
	RootCmd.AddCommand(startCMD)
}
