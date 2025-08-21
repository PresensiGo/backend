package injector

import (
	"api/internal/features/school/handlers"
	"api/internal/features/school/repositories"
	"api/internal/features/school/services"
	"api/pkg/database"
	"github.com/google/wire"
)

type SchoolHandler struct {
	School *handlers.School
}

func NewSchoolHandler(
	school *handlers.School,
) *SchoolHandler {
	return &SchoolHandler{
		School: school,
	}
}

var (
	Set = wire.NewSet(
		handlers.NewSchool,
		services.NewSchool,
		repositories.NewSchool,
		database.New,
	)
)
