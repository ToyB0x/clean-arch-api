package api

import (
	"compress/flate"
	"net/http"
	"strconv"

	cmiddleware "github.com/toaru/clean-arch-api/pkg/server/interface/api/middleware"
	"github.com/toaru/clean-arch-api/pkg/server/usecase"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/go-chi/chi"
)

type Config struct {
	ScheduleUsecase    usecase.ScheduleUsecase
	ReservationUsecase usecase.ReservationUsecase
}

type Handler struct {
	*Config
	router chi.Router
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func NewHandler(config *Config) *Handler {
	h := &Handler{
		Config: config,
		router: newRouter(),
	}

	h.router.Get("/schedules/id/{schedule_id}", h.handleGetScheduleByID)
	h.router.Get("/schedules/date/{year}/{month}", h.handleGetScheduleByMonth)
	h.router.Get("/schedules-memstore/date/{year}/{month}", h.handleGetScheduleByMonthMemStore)

	// make reservation
	h.router.Post("/reservations", h.handlePostReservation)

	return h
}

func newRouter() chi.Router {
	r := chi.NewRouter()
	corsOptions := cors.Options{
		AllowedOrigins: []string{"*"}, // for development
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "JWT"},
		MaxAge:         3600, // Maximum value not ignored by any of major browsers
	}
	r.Use(cors.New(corsOptions).Handler)
	r.Use(middleware.Logger)
	r.Use(cmiddleware.LimitSizeByContentLengthHeader)
	r.Use(cmiddleware.GetCloudFlareIp)
	r.Use(middleware.Recoverer)

	compressor := middleware.NewCompressor(flate.DefaultCompression)
	r.Use(compressor.Handler)
	return r
}

func extractDateParams(r *http.Request, strParams []string) ([]int, error) {
	arr := make([]int, len(strParams))
	for i, s := range strParams {
		p, err := strconv.Atoi(chi.URLParam(r, s))
		if err != nil {
			return arr, err
		}
		arr[i] = p
	}

	return arr, nil
}
