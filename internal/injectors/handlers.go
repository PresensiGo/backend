//go:build wireinject
// +build wireinject

package injectors

import (
	"api/database"
	"api/internal/handlers"
	"api/internal/services"
	"github.com/google/wire"
)

func InitAuthHandler() *handlers.AuthHandler {
	wire.Build(
		handlers.NewAuthHandler,
		services.NewAuthService,
		database.NewDatabase,
	)
	return nil
}

func InitBatchHandler() *handlers.BatchHandler {
	wire.Build(
		handlers.NewBatchHandler,
		services.NewBatchService,
		database.NewDatabase,
	)
	return nil
}

func InitClassHandler() *handlers.ClassHandler {
	wire.Build(
		handlers.NewClassHandler,
		services.NewClassService,
		database.NewDatabase,
	)
	return nil
}

func InitExcelHandler() *handlers.ExcelHandler {
	wire.Build(
		handlers.NewExcelHandler,
		services.NewExcelService,
		database.NewDatabase,
	)
	return nil
}

func InitMajorHandler() *handlers.MajorHandler {
	wire.Build(
		handlers.NewMajorHandler,
		services.NewMajorService,
		database.NewDatabase,
	)
	return nil
}

func InitResetService() *services.ResetService {
	wire.Build(
		services.NewResetService,
		database.NewDatabase,
	)
	return nil
}
