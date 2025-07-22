//go:build wireinject
// +build wireinject

package injectors

import (
	"api/database"
	"api/internal/handlers"
	"api/internal/services"
	"github.com/google/wire"
)

func InitAuthHandler() *handlers.Auth {
	wire.Build(
		handlers.NewAuth,
		services.NewAuth,
		database.NewDatabase,
	)
	return nil
}

func InitBatchHandler() *handlers.Batch {
	wire.Build(
		handlers.NewBatch,
		services.NewBatch,
		database.NewDatabase,
	)
	return nil
}

func InitClassHandler() *handlers.Class {
	wire.Build(
		handlers.NewClass,
		services.NewClass,
		database.NewDatabase,
	)
	return nil
}

func InitExcelHandler() *handlers.Excel {
	wire.Build(
		handlers.NewExcel,
		services.NewExcel,
		database.NewDatabase,
	)
	return nil
}

func InitMajorHandler() *handlers.Major {
	wire.Build(
		handlers.NewMajor,
		services.NewMajor,
		database.NewDatabase,
	)
	return nil
}

func InitResetService() *services.Reset {
	wire.Build(
		services.NewReset,
		database.NewDatabase,
	)
	return nil
}

func InitStudentHandler() *handlers.Student {
	wire.Build(
		handlers.NewStudent,
		services.NewStudent,
		database.NewDatabase,
	)
	return nil
}
