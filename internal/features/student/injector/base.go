package injector

import (
	repositories2 "api/internal/features/classroom/repositories"
	"api/internal/features/major/repositories"
	"api/internal/features/student/handlers"
	repositories3 "api/internal/features/student/repositories"
	"api/internal/features/student/services"
	"api/pkg/database"
	"github.com/google/wire"
)

type StudentHandlers struct {
	Student *handlers.Student
}

func NewStudentHandlers(student *handlers.Student) *StudentHandlers {
	return &StudentHandlers{
		Student: student,
	}
}

var (
	StudentSet = wire.NewSet(
		handlers.NewStudent,
		services.NewStudent,
		repositories.NewMajor,
		repositories2.NewClassroom,
		repositories3.NewStudent,
		database.New,
	)
)
