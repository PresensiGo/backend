package injector

import (
	"api/internal/features/classroom/handlers"
	repositories2 "api/internal/features/classroom/repositories"
	"api/internal/features/classroom/services"
	"api/internal/features/major/repositories"
	"api/pkg/database"
	"github.com/google/wire"
)

type ClassroomHandlers struct {
	Classroom *handlers.Classroom
}

func NewClassroomHandlers(classroom *handlers.Classroom) *ClassroomHandlers {
	return &ClassroomHandlers{
		Classroom: classroom,
	}
}

var (
	ClassroomSet = wire.NewSet(
		handlers.NewClassroom,
		services.NewClassroom,
		repositories.NewMajor,
		repositories2.NewClassroom,
		database.New,
	)
)
