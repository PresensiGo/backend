//go:build wireinject
// +build wireinject

package injectors

import (
	"api/database"
	"api/internal/handlers"
	"api/internal/services"
	"github.com/google/wire"
)

func InitExcelHandler() *handlers.ExcelHandler {
	wire.Build(
		handlers.NewExcelHandler,
		services.NewExcelService,
		database.NewDatabase,
	)
	return nil
}
