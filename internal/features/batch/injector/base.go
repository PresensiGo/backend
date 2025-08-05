package injector

import (
	"api/internal/features/batch/handlers"
	"api/internal/features/batch/repositories"
	"api/internal/features/batch/services"
	repositories3 "api/internal/features/classroom/repositories"
	repositories2 "api/internal/features/major/repositories"
	"api/pkg/database"
	"github.com/google/wire"
)

type BatchHandlers struct {
	Batch *handlers.Batch
}

func NewBatchHandlers(batch *handlers.Batch) *BatchHandlers {
	return &BatchHandlers{
		Batch: batch,
	}
}

var (
	BatchSet = wire.NewSet(
		handlers.NewBatch,

		services.NewBatch,

		repositories.NewBatch,
		repositories2.NewMajor,
		repositories3.NewClassroom,

		database.New,
	)
)
