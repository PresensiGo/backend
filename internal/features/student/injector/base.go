package injector

import (
	classroomRepo "api/internal/features/classroom/repositories"
	majorRepo "api/internal/features/major/repositories"
	schoolRepo "api/internal/features/school/repositories"
	"api/internal/features/student/handlers"
	"api/internal/features/student/repositories"
	"api/internal/features/student/services"
	"api/pkg/database"
	"github.com/google/wire"
)

type StudentHandlers struct {
	Student     *handlers.Student
	StudentAuth *handlers.StudentAuth
}

func NewStudentHandlers(
	student *handlers.Student, studentAuth *handlers.StudentAuth,
) *StudentHandlers {
	return &StudentHandlers{
		Student:     student,
		StudentAuth: studentAuth,
	}
}

var (
	StudentSet = wire.NewSet(
		handlers.NewStudent,
		handlers.NewStudentAuth,

		services.NewStudent,
		services.NewStudentAuth,

		majorRepo.NewMajor,
		classroomRepo.NewClassroom,
		repositories.NewStudent,
		repositories.NewStudentToken,
		schoolRepo.NewSchool,

		database.New,
	)
)
