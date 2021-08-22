package api_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/toaru/clean-arch-api/pkg/server/interface/api"

	"github.com/toaru/clean-arch-api/pkg/server/domain/model"

	"github.com/toaru/clean-arch-api/pkg/server/usecase/mock"

	"github.com/go-chi/chi"
)

func TestGetScheduleByMonth(t *testing.T) {
	testRequest2020_1, _ := http.NewRequest("GET", "/schedules/date/2020/1", nil)
	testRequestNoParam, _ := http.NewRequest("GET", "/schedules/date/", nil)
	// test data
	schedules := []model.Schedule{
		{
			ID:           "1",
			MaxAvailable: 10,
			Stock:        1,
			ScheduleDate: model.ScheduleDate{
				Year:  2020,
				Month: 1,
				Day:   1,
				Hour:  9,
				Min:   20,
			},
		},
	}

	type fields struct {
		Config *api.Config
		router chi.Router
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
		wantBody string
	}{
		{
			"params 2020/1 / no exist data",
			fields{
				router: chi.NewRouter(),
				Config: &api.Config{
					ScheduleUsecase: &mock.ScheduleUsecase{
						OnGetByMonth: func(ctx context.Context, year, month int) ([]model.Schedule, error) {
							return nil, nil
						},
					},
				},
			},
			args{r: testRequest2020_1},
			200,
			`null`,
		},
		{
			"params 2020/1 / exist data",
			fields{
				router: chi.NewRouter(),
				Config: &api.Config{
					ScheduleUsecase: &mock.ScheduleUsecase{
						OnGetByMonth: func(ctx context.Context, year, month int) ([]model.Schedule, error) {
							return schedules, nil
						},
					},
				},
			},
			args{r: testRequest2020_1},
			200,
			"[{\"id\":\"1\",\"max_available\":10,\"stock\":1,\"year\":2020,\"month\":1,\"day\":1,\"hour\":9,\"min\":20}]",
		},
		{
			"no params",
			fields{
				router: chi.NewRouter(),
				Config: &api.Config{
					ScheduleUsecase: &mock.ScheduleUsecase{
						OnGetByMonth: func(ctx context.Context, year, month int) ([]model.Schedule, error) {
							return nil, nil
						},
					},
				},
			},
			args{r: testRequestNoParam},
			404,
			`404 page not found`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := api.NewHandler(tt.fields.Config)
			w := httptest.NewRecorder()
			h.ServeHTTP(w, tt.args.r)

			if tt.wantCode != w.Code {
				t.Fatalf("test %q: want status code %d got %d", tt.name, tt.wantCode, w.Code)
			}

			if tt.wantBody != strings.TrimSpace(w.Body.String()) {
				t.Fatalf("test %q: want response body %q got %q", tt.name, tt.wantBody, w.Body)
			}
		})
	}
}

func TestGetScheduleByMonthMemStore(t *testing.T) {
	testRequest2020_1, _ := http.NewRequest("GET", "/schedules-memstore/date/2020/1", nil)
	testRequestNoParam, _ := http.NewRequest("GET", "/schedules-memstore/date/", nil)
	// test data
	schedules := []model.Schedule{
		{
			ID:           "1",
			MaxAvailable: 10,
			Stock:        1,
			ScheduleDate: model.ScheduleDate{
				Year:  2020,
				Month: 1,
				Day:   1,
				Hour:  9,
				Min:   20,
			},
		},
	}

	type fields struct {
		Config *api.Config
		router chi.Router
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
		wantBody string
	}{
		{
			"params 2020/1 / no exist data",
			fields{
				router: chi.NewRouter(),
				Config: &api.Config{
					ScheduleUsecase: &mock.ScheduleUsecase{
						OnGetByMonthMemStore: func(ctx context.Context, year, month int) ([]model.Schedule, error) {
							return nil, nil
						},
					},
				},
			},
			args{r: testRequest2020_1},
			200,
			`null`,
		},
		{
			"params 2020/1 / exist data",
			fields{
				router: chi.NewRouter(),
				Config: &api.Config{
					ScheduleUsecase: &mock.ScheduleUsecase{
						OnGetByMonthMemStore: func(ctx context.Context, year, month int) ([]model.Schedule, error) {
							return schedules, nil
						},
					},
				},
			},
			args{r: testRequest2020_1},
			200,
			"[{\"id\":\"1\",\"max_available\":10,\"stock\":1,\"year\":2020,\"month\":1,\"day\":1,\"hour\":9,\"min\":20}]",
		},
		{
			"no params",
			fields{
				router: chi.NewRouter(),
				Config: &api.Config{
					ScheduleUsecase: &mock.ScheduleUsecase{
						OnGetByMonth: func(ctx context.Context, year, month int) ([]model.Schedule, error) {
							return nil, nil
						},
					},
				},
			},
			args{r: testRequestNoParam},
			404,
			`404 page not found`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := api.NewHandler(tt.fields.Config)
			w := httptest.NewRecorder()
			h.ServeHTTP(w, tt.args.r)

			if tt.wantCode != w.Code {
				t.Fatalf("test %q: want status code %d got %d", tt.name, tt.wantCode, w.Code)
			}

			if tt.wantBody != strings.TrimSpace(w.Body.String()) {
				t.Fatalf("test %q: want response body %q got %q", tt.name, tt.wantBody, w.Body)
			}
		})
	}
}

func TestGetScheduleByID(t *testing.T) {
	testRequest2020_1_1_1_0, _ := http.NewRequest("GET", "/schedules/id/1", nil)

	// test data
	schedules := []model.Schedule{
		{
			ID:           "1",
			MaxAvailable: 10,
			Stock:        1,
			ScheduleDate: model.ScheduleDate{
				Year:  2020,
				Month: 1,
				Day:   1,
				Hour:  9,
				Min:   20,
			},
		},
	}

	type fields struct {
		Config *api.Config
		router chi.Router
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
		wantBody string
	}{
		{
			"params 2020/1/1/1/0 / no exist data",
			fields{
				router: chi.NewRouter(),
				Config: &api.Config{
					ScheduleUsecase: &mock.ScheduleUsecase{
						OnGetByID: func(ctx context.Context, scheduleID string) (*model.Schedule, error) {
							return nil, nil
						},
					},
				},
			},
			args{r: testRequest2020_1_1_1_0},
			200,
			`null`,
		},
		{
			"params 2020/1/1/1/0 / exist data",
			fields{
				router: chi.NewRouter(),
				Config: &api.Config{
					ScheduleUsecase: &mock.ScheduleUsecase{
						OnGetByID: func(ctx context.Context, scheduleID string) (*model.Schedule, error) {
							return &schedules[0], nil
						},
					},
				},
			},
			args{r: testRequest2020_1_1_1_0},
			200,
			"{\"id\":\"1\",\"max_available\":10,\"stock\":1,\"year\":2020,\"month\":1,\"day\":1,\"hour\":9,\"min\":20}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := api.NewHandler(tt.fields.Config)
			w := httptest.NewRecorder()
			h.ServeHTTP(w, tt.args.r)

			if tt.wantCode != w.Code {
				t.Fatalf("test %q: want status code %d got %d", tt.name, tt.wantCode, w.Code)
			}

			if tt.wantBody != strings.TrimSpace(w.Body.String()) {
				t.Fatalf("test %q: want response body %q got %q", tt.name, tt.wantBody, w.Body)
			}
		})
	}
}
