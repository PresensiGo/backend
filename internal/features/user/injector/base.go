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
	User  *handlers.User
}

func NewUserHandlers(
	auth *handlers.Auth, admin *handlers.Admin, user *handlers.User,
) *UserHandlers {
	return &UserHandlers{
		Auth:  auth,
		Admin: admin,
		User:  user,
	}
}

var (
	UserSet = wire.NewSet(
		handlers.NewAuth,
		handlers.NewAdmin,
		handlers.NewUser,

		services.NewAuth,
		services.NewAdmin,
		services.NewUser,

		repositories.NewUser,
		repositories.NewUserToken,
		repositories2.NewSchool,

		database.New,
	)
)
