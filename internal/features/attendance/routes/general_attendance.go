package routes

import (
	"api/internal/features/attendance/handlers"
	"api/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterGeneralAttendance(g *gin.RouterGroup, handler *handlers.GeneralAttendance) {
	relativePath := "/general_attendance"
	{
		group := g.Group(relativePath).Use(middlewares.AuthMiddleware())

		group.POST("", handler.Create)
		group.GET("", handler.GetAll)
		group.GET("/:general_attendance_id", handler.Get)
		group.PUT("/:general_attendance_id", handler.Update)
		group.DELETE("/:general_attendance_id", handler.Delete)
	}

	{
		group := g.Group("/general-attendances").Use(middlewares.StudentMiddleware())

		group.POST("/records/student", handler.CreateRecordStudent)
	}
}
