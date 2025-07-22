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
