package mysql_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/toaru/clean-arch-api/cmd"

	"github.com/toaru/clean-arch-api/pkg/server/infra/store/mysql"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/toaru/clean-arch-api/pkg/server/infra/store/mysql/models"

	"github.com/toaru/clean-arch-api/pkg/server/domain/model"
	"github.com/toaru/clean-arch-api/pkg/utils"

	"github.com/toaru/clean-arch-api/pkg/server/infra/store"

	"github.com/toaru/clean-arch-api/pkg/server/domain/repository"
)

func init() {
	cmd.GetConfigs("../../../../../config")
}

func TestNewReservationRepository(t *testing.T) {
	type args struct {
		sqlHandler store.SqlHandler
	}
	tests := []struct {
		name string
		args args
		want repository.ReservationRepository
	}{
		{
			"success",
			args{*store.NewSqlHandler("test")},
			&mysql.ReservationStore{SqlHandler: *store.NewSqlHandler("test")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mysql.NewReservationRepository(tt.args.sqlHandler); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("NewReservationRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReservationStore_CreatIfAvailable(t *testing.T) {
	err := utils.MigrateDrop("schema")
	if err != nil {
		panic(err)
	}
	err = utils.MigrateUp("schema")
	if err != nil {
		panic(err)
	}

	// insert test data
	dummySchedules := []model.Schedule{
		{
			ID:           "1",
			MaxAvailable: 1,
			Stock:        1,
			ScheduleDate: model.ScheduleDate{
				Year:  2021,
				Month: 1,
				Day:   1,
				Hour:  0,
				Min:   0,
			},
		},
		{
			ID:           "2",
			MaxAvailable: 2,
			Stock:        2,
			ScheduleDate: model.ScheduleDate{
				Year:  2021,
				Month: 1,
				Day:   2,
				Hour:  0,
				Min:   0,
			},
		},
	}
	sf := &mysql.ScheduleStore{SqlHandler: *store.NewSqlHandler("test")}
	for _, dummySchedule := range dummySchedules {
		err = sf.Save(context.Background(), &dummySchedule)
		if err != nil {
			panic(err)
		}
	}

	// insert test data
	dummyReservation := models.Reservation{
		ID:         "uuid1",
		ScheduleID: "1",
		TicketID:   "ticket1",
	}

	err = dummyReservation.Insert(context.Background(), store.NewSqlHandler("test").Conn, boil.Infer())
	if err != nil {
		panic(err)
	}

	type fields struct {
		SqlHandler store.SqlHandler
	}
	type args struct {
		ctx         context.Context
		reservation *model.Reservation
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"successfully create",
			fields{SqlHandler: *store.NewSqlHandler("test")},
			args{
				ctx: context.Background(),
				reservation: &model.Reservation{
					ID:         "uuid-1",
					ScheduleID: "1",
					TicketID:   "ticket-1",
				},
			},
			false,
		},
		{
			"not available",
			fields{SqlHandler: *store.NewSqlHandler("test")},
			args{
				ctx: context.Background(),
				reservation: &model.Reservation{
					ID:         "uuid-1",
					ScheduleID: "1",
					TicketID:   "ticket-1",
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &mysql.ReservationStore{
				SqlHandler: tt.fields.SqlHandler,
			}
			if err := s.CreatIfAvailable(tt.args.ctx, tt.args.reservation); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
