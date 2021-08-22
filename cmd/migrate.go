package cmd

import (
	"github.com/toaru/clean-arch-api/config"
	"github.com/toaru/clean-arch-api/pkg/utils"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate DB",
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.MigrateUp(config.Configs.MIGRATION_FILE_DIR)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
