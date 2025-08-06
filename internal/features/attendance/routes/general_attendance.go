package routes

import (
	"api/internal/features/attendance/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterGeneralAttendance(g *gin.RouterGroup, handler *handlers.GeneralAttendance) {
	group := g.Group("/general_attendances")

	group.POST("", handler.Create)
	group.GET("", handler.GetAll)
	group.PUT("/:general_attendance_id", handler.Update)
	group.DELETE("/:general_attendance_id", handler.Delete)
}
