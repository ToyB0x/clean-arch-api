package cmd

import (
	"github.com/toaru/clean-arch-api/pkg/utils"

	"github.com/spf13/cobra"
)

var RESERVATIONS_PER_Schedule int
var insertReservationCmd = &cobra.Command{
	Use:   "insert-reservations",
	Short: "Bulk insert test reservations",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := utils.BulkInsertTestReservations(RESERVATIONS_PER_Schedule)
		return err
	},
}

func init() {
	rootCmd.AddCommand(insertReservationCmd)
	insertReservationCmd.Flags().IntVarP(&RESERVATIONS_PER_Schedule, "number", "n", 100, "specify insert reservation number per day")
}
