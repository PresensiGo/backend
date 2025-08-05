package injector

import (
	repositories5 "api/internal/features/attendance/repositories"
	"api/internal/features/batch/repositories"
	repositories3 "api/internal/features/classroom/repositories"
	"api/internal/features/data/handlers"
	"api/internal/features/data/services"
	repositories2 "api/internal/features/major/repositories"
	repositories4 "api/internal/features/student/repositories"
	"api/pkg/database"
	"github.com/google/wire"
)

type DataHandlers struct {
	Excel *handlers.Excel
	Reset *handlers.Reset
}

func NewDataHandlers(excel *handlers.Excel, reset *handlers.Reset) *DataHandlers {
	return &DataHandlers{
		Excel: excel,
		Reset: reset,
	}
}

var (
	DataSet = wire.NewSet(
		handlers.NewExcel,
		handlers.NewReset,

		services.NewExcel,
		services.NewReset,

		repositories.NewBatch,
		repositories2.NewMajor,
		repositories3.NewClassroom,
		repositories4.NewStudent,
		repositories5.NewLateness,

		database.New,
	)
)
