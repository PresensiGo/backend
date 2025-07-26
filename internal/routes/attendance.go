package routes

import (
	"api/internal/injectors"
	"github.com/gin-gonic/gin"
)

func RegisterAttendance(g *gin.RouterGroup) {
	group := g.Group("/attendances")
	handler := injectors.InitAttendanceHandler()

	group.POST("", handler.Create)
	group.GET("/classrooms/:classroom_id", handler.GetAll)
	group.DELETE("/:attendance_id", handler.Delete)
}
