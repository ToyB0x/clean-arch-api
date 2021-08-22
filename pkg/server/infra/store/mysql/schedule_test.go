package mysql_test

import (
	"context"
	"log"
	"reflect"
	"testing"

	"github.com/toaru/clean-arch-api/pkg/server/domain/model"
	"github.com/toaru/clean-arch-api/pkg/utils"

	"github.com/toaru/clean-arch-api/cmd"
	"github.com/toaru/clean-arch-api/pkg/server/infra/store/mysql"

	"github.com/toaru/clean-arch-api/pkg/server/infra/store"

	"github.com/toaru/clean-arch-api/pkg/server/domain/repository"
)

func init() {
	cmd.GetConfigs("../../../../../config")
}

func TestNewScheduleRepository(t *testing.T) {
	type args struct {
		sqlHandler store.SqlHandler
	}
	tests := []struct {
		name string
		args args
		want repository.ScheduleRepository
	}{
		{
			"success",
			args{*store.NewSqlHandler("test")},
			&mysql.ScheduleStore{SqlHandler: *store.NewSqlHandler("test")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mysql.NewScheduleRepository(tt.args.sqlHandler); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("NewScheduleRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScheduleStore_Save(t *testing.T) {
	err := utils.MigrateDrop("schema")
	if err != nil {
		log.Println(err)
	}
	err = utils.MigrateUp("schema")
	if err != nil {
		log.Println(err)
	}

	type fields struct {
		SqlHandler store.SqlHandler
	}
	type args struct {
		ctx      context.Context
		schedule *model.Schedule
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
				schedule: &model.Schedule{
					MaxAvailable: 10,
					ScheduleDate: model.ScheduleDate{
						Year:  2020,
						Month: 1,
						Day:   2,
						Hour:  9,
						Min:   20,
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &mysql.ScheduleStore{
				SqlHandler: tt.fields.SqlHandler,
			}
			if err := s.Save(tt.args.ctx, tt.args.schedule); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestScheduleStore_FindByMonth(t *testing.T) {
	err := utils.MigrateDrop("schema")
	if err != nil {
		panic(err)
	}
	err = utils.MigrateUp("schema")
	if err != nil {
		panic(err)
	}

	dummySchedules := []model.Schedule{
		{
			ID:           "1",
			MaxAvailable: 10,
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
			MaxAvailable: 10,
			Stock:        2,
			ScheduleDate: model.ScheduleDate{
				Year:  2021,
				Month: 1,
				Day:   2,
				Hour:  9,
				Min:   20,
			},
		},
		{
			ID:           "3",
			MaxAvailable: 10,
			Stock:        3,
			ScheduleDate: model.ScheduleDate{
				Year:  2021,
				Month: 2,
				Day:   1,
				Hour:  0,
				Min:   0,
			},
		},
	}

	type fields struct {
		SqlHandler store.SqlHandler
	}
	type args struct {
		ctx         context.Context
		year, month int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Schedule
		wantErr bool
	}{
		{
			"successfully find month 1",
			fields{SqlHandler: *store.NewSqlHandler("test")},
			args{
				ctx:   context.Background(),
				year:  2021,
				month: 1,
			},
			dummySchedules[0:2],
			false,
		},
		{
			"successfully find month 2",
			fields{SqlHandler: *store.NewSqlHandler("test")},
			args{
				ctx:   context.Background(),
				year:  2021,
				month: 2,
			},
			dummySchedules[2:3],
			false,
		},
		{
			"no hit month 3",
			fields{SqlHandler: *store.NewSqlHandler("test")},
			args{
				ctx:   context.Background(),
				year:  2021,
				month: 3,
			},
			[]model.Schedule{},
			false,
		},
	}

	// insert test data
	s := &mysql.ScheduleStore{SqlHandler: *store.NewSqlHandler("test")}
	for _, dummySchedule := range dummySchedules {
		err = s.Save(context.Background(), &dummySchedule)
		if err != nil {
			panic(err)
		}
	}

	for _, tt := range tests {
		s := &mysql.ScheduleStore{
			SqlHandler: tt.fields.SqlHandler,
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.FindByMonth(tt.args.ctx, tt.args.year, tt.args.month)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScheduleStore_FindByDateTime(t *testing.T) {
	err := utils.MigrateDrop("schema")
	if err != nil {
		panic(err)
	}
	err = utils.MigrateUp("schema")
	if err != nil {
		panic(err)
	}

	dummySchedules := []model.Schedule{
		{
			ID:           "1",
			MaxAvailable: 10,
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
			MaxAvailable: 10,
			Stock:        2,
			ScheduleDate: model.ScheduleDate{
				Year:  2021,
				Month: 1,
				Day:   1,
				Hour:  1,
				Min:   1,
			},
		},
	}

	type fields struct {
		SqlHandler store.SqlHandler
	}
	type args struct {
		ctx        context.Context
		scheduleID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Schedule
		wantErr bool
	}{
		{
			"successfully find id 1",
			fields{SqlHandler: *store.NewSqlHandler("test")},
			args{
				ctx:        context.Background(),
				scheduleID: "1",
			},
			&dummySchedules[0],
			false,
		},
		{
			"successfully find id 2",
			fields{SqlHandler: *store.NewSqlHandler("test")},
			args{
				ctx:        context.Background(),
				scheduleID: "2",
			},
			&dummySchedules[1],
			false,
		},
		{
			"no hit id 3",
			fields{SqlHandler: *store.NewSqlHandler("test")},
			args{
				ctx:        context.Background(),
				scheduleID: "3",
			},
			nil,
			true,
		},
	}

	// insert test data
	s := &mysql.ScheduleStore{SqlHandler: *store.NewSqlHandler("test")}
	for _, dummySchedule := range dummySchedules {
		err = s.Save(context.Background(), &dummySchedule)
		if err != nil {
			panic(err)
		}
	}

	for _, tt := range tests {
		s := &mysql.ScheduleStore{
			SqlHandler: tt.fields.SqlHandler,
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.FindByID(tt.args.ctx, tt.args.scheduleID)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
