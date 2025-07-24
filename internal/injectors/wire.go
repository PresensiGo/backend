//go:build wireinject
// +build wireinject

package injectors

import (
	"api/internal/handlers"
	"api/internal/repository"
	"api/internal/services"
	"api/pkg/database"
	"github.com/google/wire"
)

func InitAuthHandler() *handlers.Auth {
	wire.Build(
		handlers.NewAuth,
		services.NewAuth,
		database.New,
	)
	return nil
}

func InitBatchHandler() *handlers.Batch {
	wire.Build(
		handlers.NewBatch,
		services.NewBatch,
		database.New,
	)
	return nil
}

func InitClassHandler() *handlers.Class {
	wire.Build(
		handlers.NewClass,
		services.NewClass,
		database.New,
	)
	return nil
}

func InitClassMajorHandler() *handlers.ClassMajor {
	wire.Build(
		handlers.NewClassMajor,
		services.NewClassMajor,
		repository.NewClass,
		repository.NewMajor,
		database.New,
	)
	return nil
}

func InitExcelHandler() *handlers.Excel {
	wire.Build(
		handlers.NewExcel,
		services.NewExcel,
		database.New,
	)
	return nil
}

func InitMajorHandler() *handlers.Major {
	wire.Build(
		handlers.NewMajor,
		services.NewMajor,
		database.New,
	)
	return nil
}

func InitResetService() *services.Reset {
	wire.Build(
		services.NewReset,
		database.New,
	)
	return nil
}

func InitStudentHandler() *handlers.Student {
	wire.Build(
		handlers.NewStudent,
		services.NewStudent,
		repository.NewStudent,
		database.New,
	)
	return nil
}
