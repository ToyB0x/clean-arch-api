package cmd

import (
	"time"

	"github.com/toaru/clean-arch-api/pkg/utils"

	"github.com/spf13/cobra"
)

var insertScheduleCmd = &cobra.Command{
	Use:   "insert-schedules",
	Short: "Bulk insert test schedules",
	RunE: func(cmd *cobra.Command, args []string) error {
		startDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
		totalDays := 365 * 3
		hours := []int{8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
		mins := []int{0, 20, 40}
		maxAvailable := 300
		err := utils.BulkInsertSchedules(startDate, totalDays, hours, mins, maxAvailable)
		return err
	},
}

func init() {
	rootCmd.AddCommand(insertScheduleCmd)
}
