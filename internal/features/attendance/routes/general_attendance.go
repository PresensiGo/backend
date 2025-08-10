package routes

import (
	"api/internal/features/attendance/handlers"
	"api/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterGeneralAttendance(g *gin.RouterGroup, handler *handlers.GeneralAttendance) {
	group := g.Group("/general_attendances").
		Use(middlewares.AuthMiddleware())

	group.POST("", handler.Create)
	group.GET("", handler.GetAll)
	group.GET("/:general_attendance_id", handler.Get)
	group.PUT("/:general_attendance_id", handler.Update)
	group.DELETE("/:general_attendance_id", handler.Delete)

	{

		studentGroup := g.Group("/general-attendances").
			Use(middlewares.StudentMiddleware())

		studentGroup.POST("/records/student", handler.CreateGeneralAttendanceRecordStudent)
	}
}
