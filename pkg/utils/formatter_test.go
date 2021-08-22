package utils_test

import (
	"testing"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/toaru/clean-arch-api/pkg/utils"
)

func TestDateFormatter(t *testing.T) {
	type args struct {
		year, month, day, hour, min uint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"success",
			args{2000, 1, 1, 1, 1},
			"2000-01/01-01:01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.DateFormatter(tt.args.year, tt.args.month, tt.args.day, tt.args.hour, tt.args.min); got != tt.want {
				t.Errorf("NewScheduleUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
