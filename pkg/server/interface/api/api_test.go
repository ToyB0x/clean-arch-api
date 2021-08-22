package api_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/toaru/clean-arch-api/pkg/server/interface/api"

	"github.com/toaru/clean-arch-api/pkg/server/usecase/mock"
)

func TestNewHandler(t *testing.T) {
	type args struct {
		config *api.Config
	}
	tests := []struct {
		name string
		args args
		want *api.Handler
	}{
		{
			"success",
			args{
				&api.Config{
					ReservationUsecase: &mock.ReservationUsecase{},
				},
			},
			&api.Handler{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := api.NewHandler(tt.args.config); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractDateParams(t *testing.T) {
	testRequest2020_1, _ := http.NewRequest("GET", "/reservations/2020-1", nil)
	testRequest2020_1_1_0_0, _ := http.NewRequest("GET", "/reservations/2020-1-1-0-0", nil)

	type args struct {
		r      *http.Request
		params []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Month",
			args{
				r:      testRequest2020_1,
				params: []string{"2020", "1"},
			},
			[]int{2020, 1},
		},
		{
			"DateTime",
			args{
				r:      testRequest2020_1_1_0_0,
				params: []string{"2020", "1", "1", "0", "0"},
			},
			[]int{2020, 1, 1, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := api.ExtractDateParams(tt.args.r, tt.args.params); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
