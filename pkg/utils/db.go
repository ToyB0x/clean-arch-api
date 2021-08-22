package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4/database/mysql"

	"github.com/toaru/clean-arch-api/config"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/toaru/clean-arch-api/pkg/server/infra/store"
	"github.com/volatiletech/sqlboiler/queries"

	"github.com/toaru/clean-arch-api/pkg/server/infra/store/mysql/models"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateUp(filePath string) error {
	m, err := getMigrateInstance("file://" + filePath)
	if err != nil {
		return err
	}
	return m.Up()
}

func MigrateDrop(filePath string) error {
	m, err := getMigrateInstance("file://" + filePath)
	if err != nil {
		return err
	}
	return m.Drop()
}

func getMigrateInstance(filePath string) (*migrate.Migrate, error) {
	con := store.NewSqlHandler(config.Configs.APP_ENV).Conn
	driver, _ := mysql.WithInstance(con, &mysql.Config{})
	return migrate.NewWithDatabaseInstance(
		filePath,
		"clean-arch-api",
		driver,
	)
}

func BulkInsertSchedules(startDate time.Time, totalDays int, hours, mins []int, maxAvailable int) error {
	rowNum := totalDays * len(hours) * len(mins)
	fmt.Println(len(hours)*len(mins)*maxAvailable, "reservations per day")
	fmt.Println(rowNum*maxAvailable, "reservations available totally")
	fmt.Println("creating", rowNum, "rows")

	con := store.NewSqlHandler(config.Configs.APP_ENV).Conn
	dbName := config.Configs.DB_NAME
	tableName := "schedule"
	col1 := "id"
	col2 := "date"
	col3 := "hour"
	col4 := "min"
	col5 := "max_available"
	col6 := "stock"
	col7 := "created_at"
	col8 := "updated_at"

	bulkInsertTriggerCount := 10000
	bulkInsertedCount := 0
	var schedules []models.Schedule
	for i := 0; i <= totalDays; i++ {
		date := startDate.AddDate(0, 0, i)
		for j, hour := range hours {
			for k, min := range mins {
				schedule := models.Schedule{
					ID:           fmt.Sprintf("%d-%d-%d", i, j, k),
					Date:         time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC),
					Hour:         uint8(hour),
					Min:          uint8(min),
					MaxAvailable: uint(maxAvailable),
				}
				schedules = append(schedules, schedule)

				// bulk insret
				isAll := len(schedules) == rowNum
				isTrigerSize := len(schedules) >= bulkInsertTriggerCount
				if isAll || isTrigerSize {
					queryStrBefore := fmt.Sprintf("INSERT INTO %s.%s (%s, %s, %s, %s, %s, %s, %s, %s)", dbName, tableName, col1, col2, col3, col4, col5, col6, col7, col8)
					queryStrMiddle := ` VALUES `
					queryStrAfter := `;`

					for _, s := range schedules {
						d := fmt.Sprintf("%04d%02d%02d", s.Date.Year(), s.Date.Month(), s.Date.Day())
						createdAt := time.Now().Format("2006-01-02 15:04:05")
						updatedAt := time.Now().Format("2006-01-02 15:04:05")
						//index := fmt.Sprintf("%d-%d-%d", i, j, k)
						queryStrMiddle += fmt.Sprintf("('%s', '%s', %d, %d, %d, %d, '%s', '%s'),", s.ID, d, s.Hour, s.Min, s.MaxAvailable, s.MaxAvailable, createdAt, updatedAt)
					}
					queryStrMiddle = queryStrMiddle[:(len(queryStrMiddle) - 1)] // replace "," to ""
					queryStr := queryStrBefore + queryStrMiddle + queryStrAfter
					if err := queries.Raw(queryStr).Bind(context.Background(), con, &[]models.Schedule{}); err != nil {
						return err
					}
					bulkInsertedCount += 1
					schedules = []models.Schedule{} // clear memory
					log.Println("bulk inserted", bulkInsertedCount)
				}
			}
		}
	}
	return nil
}

func BulkInsertTestReservations(reservationPerSchedule int) error {
	con := store.NewSqlHandler(config.Configs.APP_ENV).Conn
	schedules, err := models.Schedules().All(context.Background(), con)
	if err != nil {
		return err
	}

	fmt.Println(len(schedules)*reservationPerSchedule, "reservations will be inserted")
	for i, s := range schedules {
		s.Stock = s.MaxAvailable - uint(reservationPerSchedule)
		fmt.Println("inserting", i, "/", len(schedules))
		if _, err = s.Update(context.Background(), con, boil.Whitelist(models.ScheduleColumns.Stock)); err != nil {
			return err
		}

		dbName := config.Configs.DB_NAME
		tableName := "reservation"
		col1 := "id"
		col2 := "ticket_id"
		col3 := "schedule_id"
		col4 := "created_at"
		queryStrBefore := fmt.Sprintf("INSERT INTO %s.%s (%s, %s, %s, %s)", dbName, tableName, col1, col2, col3, col4)
		queryStrMiddle := ` VALUES `
		queryStrAfter := `;`
		for j := 0; j < reservationPerSchedule; j++ {
			id := fmt.Sprintf("schedule_id_%s__loop_%d", s.ID, j)
			ticketID := "test-ticket"
			createdAt := time.Now().Format("2006-01-02 15:04:05")
			queryStrMiddle += fmt.Sprintf("('%s', '%s', '%s', '%s'),", id, ticketID, s.ID, createdAt)
		}
		queryStrMiddle = queryStrMiddle[:(len(queryStrMiddle) - 1)] // replace "," to ""
		queryStr := queryStrBefore + queryStrMiddle + queryStrAfter
		if err := queries.Raw(queryStr).Bind(context.Background(), con, &[]models.Reservation{}); err != nil {
			return err
		}
	}

	return nil
}
