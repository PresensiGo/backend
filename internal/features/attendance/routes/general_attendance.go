package routes

import (
	"api/internal/features/attendance/handlers"
	"api/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterGeneralAttendance(g *gin.RouterGroup, handler *handlers.GeneralAttendance) {
	relativePath := "/general-attendances"
	{
		group := g.Group(relativePath).Use(middlewares.AuthMiddleware())

		group.POST("", handler.CreateGeneralAttendance)
		group.POST("/:general_attendance_id/records", handler.CreateGeneralAttendanceRecord)

		group.GET("", handler.GetAllGeneralAttendances)
		group.GET("/:general_attendance_id/records", handler.GetAllGeneralAttendanceRecords)
		group.GET(
			"/:general_attendance_id/classrooms/:classroom_id/records",
			handler.GetAllGeneralAttendanceRecordsByClassroomId,
		)
		group.GET("/:general_attendance_id", handler.GetGeneralAttendance)

		group.PUT("/:general_attendance_id", handler.Update)

		group.DELETE("/:general_attendance_id", handler.DeleteGeneralAttendance)
		group.DELETE(
			"/:general_attendance_id/records/:record_id", handler.DeleteGeneralAttendanceRecord,
		)
	}

	{
		group := g.Group(relativePath).Use(middlewares.StudentMiddleware())

		group.POST("/records/student", handler.CreateGeneralAttendanceRecordStudent)

		group.GET("/student", handler.GetAllGeneralAttendancesStudent)
	}
}
