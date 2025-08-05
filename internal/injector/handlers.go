//go:build wireinject
// +build wireinject

package injector

import (
	"api/internal/features/attendance/injector"
	injector2 "api/internal/features/batch/injector"
	injector4 "api/internal/features/classroom/injector"
	injector6 "api/internal/features/data/injector"
	injector3 "api/internal/features/major/injector"
	injector5 "api/internal/features/student/injector"
	injector7 "api/internal/features/user/injector"
	"github.com/google/wire"
)

func InitAttendanceHandlers() *injector.AttendanceHandlers {
	wire.Build(
		injector.NewAttendanceHandlers,
		injector.AttendanceSet,
	)
	return nil
}

func InitBatchHandlers() *injector2.BatchHandlers {
	wire.Build(
		injector2.NewBatchHandlers,
		injector2.BatchSet,
	)
	return nil
}

func InitMajorHandlers() *injector3.MajorHandlers {
	wire.Build(
		injector3.NewMajorHandlers,
		injector3.MajorSet,
	)
	return nil
}

func InitClassroomHandlers() *injector4.ClassroomHandlers {
	wire.Build(
		injector4.NewClassroomHandlers,
		injector4.ClassroomSet,
	)
	return nil
}

func InitStudentHandlers() *injector5.StudentHandlers {
	wire.Build(
		injector5.NewStudentHandlers,
		injector5.StudentSet,
	)
	return nil
}

func InitDataHandlers() *injector6.DataHandlers {
	wire.Build(
		injector6.NewDataHandlers,
		injector6.DataSet,
	)
	return nil
}

func InitUserHandlers() *injector7.UserHandlers {
	wire.Build(
		injector7.NewUserHandlers,
		injector7.UserSet,
	)
	return nil
}
