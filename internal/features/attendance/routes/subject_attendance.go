package routes

import (
	"api/internal/features/attendance/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterSubjectAttendance(g *gin.RouterGroup, handler *handlers.SubjectAttendance) {
	group := g.Group("/subject_attendances")

	group.GET("", handler.GetAll)
}
