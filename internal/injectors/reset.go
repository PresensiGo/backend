//go:build wireinject
// +build wireinject

package injectors

import (
	"api/database"
	"api/internal/services"
	"github.com/google/wire"
)

func InitResetService() *services.ResetService {
	wire.Build(services.NewResetService, database.NewDatabase)
	return nil
}
