//go:build wireinject
// +build wireinject

package injectors

import (
	"api/internal/cron"
	"api/internal/handlers"
	"api/internal/repositories"
	"api/internal/services"
	"api/pkg/database"
	"github.com/google/wire"
)

func InitAttendanceHandler() *handlers.Attendance {
	wire.Build(
		handlers.NewAttendance,
		services.NewAttendance,
		repositories.NewAttendance,
		repositories.NewAttendanceStudent,
		repositories.NewStudent,
		database.New,
	)
	return nil
}

func InitAuthHandler() *handlers.Auth {
	wire.Build(
		handlers.NewAuth,
		services.NewAuth,
		repositories.NewUser,
		repositories.NewUserToken,
		repositories.NewSchool,
		database.New,
	)
	return nil
}

func InitBatchHandler() *handlers.Batch {
	wire.Build(
		handlers.NewBatch,
		services.NewBatch,
		database.New,
	)
	return nil
}

func InitClassroomHandler() *handlers.Classroom {
	wire.Build(
		handlers.NewClassroom,
		services.NewClassroom,
		repositories.NewClassroom,
		repositories.NewMajor,
		database.New,
	)
	return nil
}

func InitExcelHandler() *handlers.Excel {
	wire.Build(
		handlers.NewExcel,
		services.NewExcel,
		repositories.NewBatch,
		repositories.NewMajor,
		repositories.NewClassroom,
		repositories.NewStudent,
		database.New,
	)
	return nil
}

func InitMajorHandler() *handlers.Major {
	wire.Build(
		handlers.NewMajor,
		services.NewMajor,
		database.New,
	)
	return nil
}

func InitResetHandler() *handlers.Reset {
	wire.Build(
		handlers.NewReset,
		services.NewReset,
		repositories.NewBatch,
		database.New,
	)
	return nil
}

func InitStudentHandler() *handlers.Student {
	wire.Build(
		handlers.NewStudent,
		services.NewStudent,
		repositories.NewStudent,
		database.New,
	)
	return nil
}

func InitUserTokenCron() *cron.UserToken {
	wire.Build(
		cron.NewUserToken,
		repositories.NewUserToken,
		database.New,
	)
	return nil
}
