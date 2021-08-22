package server

import (
	"github.com/toaru/clean-arch-api/config"
	"github.com/toaru/clean-arch-api/pkg/server/infra/auth"
	"github.com/toaru/clean-arch-api/pkg/server/infra/memstore"
	"github.com/toaru/clean-arch-api/pkg/server/infra/store"
	"github.com/toaru/clean-arch-api/pkg/server/infra/store/mysql"
	"github.com/toaru/clean-arch-api/pkg/server/interface/api"
	"github.com/toaru/clean-arch-api/pkg/server/usecase"
)

func getConfig() api.Config {
	// infra
	sqlHandler := store.NewSqlHandler(config.Configs.APP_ENV)
	authService := auth.NewAuthService()
	scheduleRepository := mysql.NewScheduleRepository(*sqlHandler)
	reservationRepository := mysql.NewReservationRepository(*sqlHandler)
	memStoreService := memstore.NewMemStoreService(config.Configs.REDIS_PORT)

	// usecase
	scheduleUsecase := usecase.NewScheduleUsecase(scheduleRepository, reservationRepository, memStoreService)
	reservationUsecase := usecase.NewReservationUsecase(reservationRepository, authService)

	// interface
	return api.Config{
		ScheduleUsecase:    scheduleUsecase,
		ReservationUsecase: reservationUsecase,
	}
}
