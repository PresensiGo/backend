package injector

import (
	"api/internal/features/teacher/handlers"
	"api/internal/features/teacher/services"
	repositories2 "api/internal/features/user/repositories"
	"api/pkg/database"
	"github.com/google/wire"
)

type TeacherHandlers struct {
	Teacher *handlers.Teacher
}

func NewTeacherHandlers(teacher *handlers.Teacher) *TeacherHandlers {
	return &TeacherHandlers{
		Teacher: teacher,
	}
}

var (
	Set = wire.NewSet(
		handlers.NewTeacher,
		services.NewTeacher,
		repositories2.NewUser,
		database.New,
	)
)
