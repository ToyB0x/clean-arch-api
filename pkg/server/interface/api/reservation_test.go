package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/toaru/clean-arch-api/cmd"

	"github.com/toaru/clean-arch-api/pkg/server/interface/api"

	"github.com/toaru/clean-arch-api/pkg/server/usecase/mock"

	"github.com/go-chi/chi"
)

func init() {
	cmd.GetConfigs("../../../../config")
}

func TestHandler_handleCreateReservation(t *testing.T) {
	reservation1 := api.PostReservationJson{
		TicketID:   "ticket-1",
		ScheduleID: "1",
	}
	json1, _ := json.Marshal(reservation1)
	testRequest1, _ := http.NewRequest("POST", "/reservations", bytes.NewBuffer(json1))
	testRequest1.Header.Set("Content-Type", "application/json")

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
			"success",
			fields{
				router: chi.NewRouter(),
				Config: &api.Config{
					ReservationUsecase: &mock.ReservationUsecase{
						OnCreate: func(ctx context.Context, ticketID, scheduleID string) error {
							return nil
						},
					},
				},
			},
			args{r: testRequest1},
			200,
			`"ok"`,
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
