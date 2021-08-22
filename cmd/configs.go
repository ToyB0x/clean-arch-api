package cmd

import (
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
	"github.com/toaru/clean-arch-api/config"
)

var configsCmd = &cobra.Command{
	Use:   "configs",
	Short: "Show current configs",
	Run: func(cmd *cobra.Command, args []string) {
		pp.Print(config.Configs)
	},
}

func init() {
	rootCmd.AddCommand(configsCmd)
}
