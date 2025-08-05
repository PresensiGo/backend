//go:build wireinject
// +build wireinject

package injector

import (
	"api/internal/cron"
	"api/internal/features/user/repositories"
	"api/pkg/database"
	"github.com/google/wire"
)

func InitUserTokenCron() *cron.UserTokenCron {
	wire.Build(
		cron.NewUserTokenCron,
		repositories.NewUserToken,
		database.New,
	)
	return nil
}
