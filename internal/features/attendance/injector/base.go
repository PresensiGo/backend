package injector

import (
	"api/internal/features/attendance/handlers"
	"api/internal/features/attendance/repositories"
	"api/internal/features/attendance/services"
	repositories4 "api/internal/features/classroom/repositories"
	repositories3 "api/internal/features/major/repositories"
	repositories2 "api/internal/features/student/repositories"
	"api/pkg/database"
	"github.com/google/wire"
)

type AttendanceHandlers struct {
	Attendance *handlers.Attendance
	Lateness   *handlers.Lateness
}

func NewAttendanceHandlers(
	attendance *handlers.Attendance, lateness *handlers.Lateness,
) *AttendanceHandlers {
	return &AttendanceHandlers{
		Attendance: attendance,
		Lateness:   lateness,
	}
}

var (
	AttendanceSet = wire.NewSet(
		handlers.NewAttendance,
		handlers.NewLateness,

		services.NewAttendance,
		services.NewLateness,

		repositories.NewAttendance,
		repositories.NewAttendanceStudent,
		repositories2.NewStudent,
		repositories.NewLateness,
		repositories.NewLatenessDetail,
		repositories3.NewMajor,
		repositories4.NewClassroom,

		database.New,
	)
)
