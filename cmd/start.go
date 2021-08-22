package cmd

import (
	"github.com/spf13/cobra"
	"github.com/toaru/clean-arch-api/pkg/server"
)

var apiCmd = &cobra.Command{
	Use:   "start",
	Short: "Run server server",
	RunE: func(cmd *cobra.Command, args []string) error {
		server.RunServer()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
