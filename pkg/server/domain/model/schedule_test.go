package model_test

import (
	"reflect"
	"testing"

	"github.com/toaru/clean-arch-api/pkg/server/domain/model"
)

func TestNewSchedule(t *testing.T) {
	type args struct {
		maxAvailable uint
		stock        uint
		date         *model.ScheduleDate
	}
	tests := []struct {
		name string
		args args
		want *model.Schedule
	}{
		{
			name: "with padding date",
			args: args{
				maxAvailable: 1,
				stock:        1,
				date: &model.ScheduleDate{
					Year:  2020,
					Month: 1,
					Day:   1,
					Hour:  0,
					Min:   0,
				},
			},
			want: &model.Schedule{
				ID:           "00000000-0000-0000-0000-000000000000",
				MaxAvailable: 1,
				Stock:        1,
				ScheduleDate: model.ScheduleDate{
					Year:  2020,
					Month: 1,
					Day:   1,
					Hour:  0,
					Min:   0,
				},
			},
		},
		{
			name: "without padding date",
			args: args{
				date: &model.ScheduleDate{
					Year:  2020,
					Month: 10,
					Day:   10,
					Hour:  10,
					Min:   10,
				},
				maxAvailable: 1,
				stock:        1,
			},
			want: &model.Schedule{
				ID:           "00000000-0000-0000-0000-000000000000",
				MaxAvailable: 1,
				Stock:        1,
				ScheduleDate: model.ScheduleDate{
					Year:  2020,
					Month: 10,
					Day:   10,
					Hour:  10,
					Min:   10,
				},
			},
		},
		{
			name: "great number",
			args: args{
				date: &model.ScheduleDate{
					Year:  20200,
					Month: 100,
					Day:   100,
					Hour:  100,
					Min:   100,
				},
				maxAvailable: 1000,
				stock:        1000,
			},
			want: &model.Schedule{
				ID:           "00000000-0000-0000-0000-000000000000",
				MaxAvailable: 1000,
				Stock:        1000,
				ScheduleDate: model.ScheduleDate{
					Year:  20200,
					Month: 100,
					Day:   100,
					Hour:  100,
					Min:   100,
				},
			},
		},
		{
			name: "small number",
			args: args{
				date: &model.ScheduleDate{
					Year:  0,
					Month: 0,
					Day:   0,
					Hour:  0,
					Min:   0,
				},
				maxAvailable: 0,
				stock:        0,
			},
			want: &model.Schedule{
				ID:           "00000000-0000-0000-0000-000000000000",
				MaxAvailable: 0,
				Stock:        0,
				ScheduleDate: model.ScheduleDate{
					Year:  0,
					Month: 0,
					Day:   0,
					Hour:  0,
					Min:   0,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := model.NewSchedule(tt.args.date, tt.args.maxAvailable, mockUUIDGen{}); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}
