package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/toaru/clean-arch-api/pkg/server/infra/store"
	"github.com/toaru/clean-arch-api/pkg/server/infra/store/mysql/models"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/spf13/cobra"
)

var insertReservationBenchCmd = &cobra.Command{
	Use:   "insert-reservation-bench",
	Short: "Bulk insert test reservations",
	RunE: func(cmd *cobra.Command, args []string) error {
		con := store.NewSqlHandler("local")

		log.Println("start bench")
		startAt := time.Now().UnixNano()
		for i := 0; i < 100; i++ {
			r := models.Reservation{
				ID:         fmt.Sprintf("bench_%d_%d", i, time.Now().Nanosecond()),
				ScheduleID: fmt.Sprintf("%d-0-0", i),
				TicketID:   fmt.Sprintf("%d", i),
				CreatedAt:  time.Now(),
			}
			if err := r.Insert(context.Background(), con.Conn, boil.Infer()); err != nil {
				return err
			}
		}
		endAt := time.Now().UnixNano()
		log.Println("end bench")

		delta := endAt - startAt
		fmt.Printf("insert reservation: 100 times\n"+
			"Total time        : %d nano sec\n"+
			"1 item            : %d nano sec\n",
			delta, delta/100)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(insertReservationBenchCmd)
}
