package injector

import (
	"api/internal/features/attendance/handlers"
	"api/internal/features/attendance/repositories"
	"api/internal/features/attendance/services"
	batch "api/internal/features/batch/repositories"
	classroom "api/internal/features/classroom/repositories"
	major "api/internal/features/major/repositories"
	student "api/internal/features/student/repositories"
	subject "api/internal/features/subject/repositories"
	"api/pkg/database"
	"github.com/google/wire"
)

type AttendanceHandlers struct {
	GeneralAttendance *handlers.GeneralAttendance
	SubjectAttendance *handlers.SubjectAttendance
}

func NewAttendanceHandlers(
	generalAttendance *handlers.GeneralAttendance, subjectAttendance *handlers.SubjectAttendance,
) *AttendanceHandlers {
	return &AttendanceHandlers{
		GeneralAttendance: generalAttendance,
		SubjectAttendance: subjectAttendance,
	}
}

var (
	AttendanceSet = wire.NewSet(
		handlers.NewGeneralAttendance,
		handlers.NewSubjectAttendance,

		services.NewGeneralAttendance,
		services.NewSubjectAttendance,

		repositories.NewGeneralAttendance,
		student.NewStudent,
		batch.NewBatch,
		major.NewMajor,
		classroom.NewClassroom,
		repositories.NewSubjectAttendance,
		subject.NewSubject,
		repositories.NewGeneralAttendanceRecord,

		database.New,
	)
)
