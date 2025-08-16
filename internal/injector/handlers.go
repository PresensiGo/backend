//go:build wireinject
// +build wireinject

package injector

import (
	attendance "api/internal/features/attendance/injector"
	batch "api/internal/features/batch/injector"
	classroom "api/internal/features/classroom/injector"
	data "api/internal/features/data/injector"
	major "api/internal/features/major/injector"
	student "api/internal/features/student/injector"
	subject "api/internal/features/subject/injector"
	user "api/internal/features/user/injector"
	"github.com/google/wire"
)

func InitAttendanceHandlers() *attendance.AttendanceHandlers {
	wire.Build(
		attendance.NewAttendanceHandlers,
		attendance.AttendanceSet,
	)
	return nil
}

func InitBatchHandlers() *batch.BatchHandlers {
	wire.Build(
		batch.NewBatchHandlers,
		batch.BatchSet,
	)
	return nil
}

func InitMajorHandlers() *major.MajorHandlers {
	wire.Build(
		major.NewMajorHandlers,
		major.MajorSet,
	)
	return nil
}

func InitClassroomHandlers() *classroom.ClassroomHandlers {
	wire.Build(
		classroom.NewClassroomHandlers,
		classroom.ClassroomSet,
	)
	return nil
}

func InitStudentHandlers() *student.StudentHandlers {
	wire.Build(
		student.NewStudentHandlers,
		student.StudentSet,
	)
	return nil
}

func InitDataHandlers() *data.DataHandlers {
	wire.Build(
		data.NewDataHandlers,
		data.DataSet,
	)
	return nil
}

func InitUserHandlers() *user.UserHandlers {
	wire.Build(
		user.NewUserHandlers,
		user.UserSet,
	)
	return nil
}

func InitSubjectHandlers() *subject.SubjectHandlers {
	wire.Build(
		subject.NewSubjectHandlers,
		subject.SubjectSet,
	)
	return nil
}
