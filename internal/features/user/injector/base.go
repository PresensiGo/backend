package injector

import (
	repositories2 "api/internal/features/school/repositories"
	"api/internal/features/user/handlers"
	"api/internal/features/user/repositories"
	"api/internal/features/user/services"
	"api/pkg/database"
	"github.com/google/wire"
)

type UserHandlers struct {
	Auth  *handlers.Auth
	Admin *handlers.Admin
}

func NewUserHandlers(auth *handlers.Auth, admin *handlers.Admin) *UserHandlers {
	return &UserHandlers{
		Auth:  auth,
		Admin: admin,
	}
}

var (
	UserSet = wire.NewSet(
		handlers.NewAuth,
		handlers.NewAdmin,

		services.NewAuth,
		services.NewAdmin,

		repositories.NewUser,
		repositories.NewUserToken,
		repositories2.NewSchool,

		database.New,
	)
)
