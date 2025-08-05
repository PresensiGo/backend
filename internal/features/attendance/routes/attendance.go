package routes

import (
	"api/internal/features/attendance/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterAttendance(g *gin.RouterGroup, handler *handlers.Attendance) {
	group := g.Group("/attendances")

	group.POST("", handler.Create)
	group.GET("/classrooms/:classroom_id", handler.GetAll)
	group.GET("/:attendance_id", handler.GetById)
	group.DELETE("/:attendance_id", handler.Delete)
}
