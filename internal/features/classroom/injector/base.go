package injector

import (
	batchRepo "api/internal/features/batch/repositories"
	"api/internal/features/classroom/handlers"
	classroomRepo "api/internal/features/classroom/repositories"
	"api/internal/features/classroom/services"
	"api/internal/features/major/repositories"
	studentRepo "api/internal/features/student/repositories"
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
		batchRepo.NewBatch,
		repositories.NewMajor,
		classroomRepo.NewClassroom,
		studentRepo.NewStudent,
		database.New,
	)
)
