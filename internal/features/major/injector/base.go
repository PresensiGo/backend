package injector

import (
	repositories2 "api/internal/features/batch/repositories"
	"api/internal/features/major/handlers"
	"api/internal/features/major/repositories"
	"api/internal/features/major/services"
	"api/pkg/database"
	"github.com/google/wire"
)

type MajorHandlers struct {
	Major *handlers.Major
}

func NewMajorHandlers(major *handlers.Major) *MajorHandlers {
	return &MajorHandlers{
		Major: major,
	}
}

var (
	MajorSet = wire.NewSet(
		handlers.NewMajor,
		services.NewMajor,
		repositories2.NewBatch,
		repositories.NewMajor,
		database.New,
	)
)
