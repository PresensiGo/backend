package injector

import (
	batch "api/internal/features/batch/repositories"
	classroom "api/internal/features/classroom/repositories"
	major "api/internal/features/major/repositories"
	school "api/internal/features/school/repositories"
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

		batch.NewBatch,
		major.NewMajor,
		classroom.NewClassroom,
		repositories.NewStudent,
		repositories.NewStudentToken,
		school.NewSchool,

		database.New,
	)
)
