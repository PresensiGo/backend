package routes

import (
	"api/internal/features/attendance/handlers"
	"api/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterSubjectAttendance(g *gin.RouterGroup, handler *handlers.SubjectAttendance) {
	group := g.Group("/batches/:batch_id/majors/:major_id/classrooms/:classroom_id/subject-attendances").
		Use(middlewares.AuthMiddleware())
	group.POST("", handler.Create)
	group.GET("", handler.GetAll)
	group.GET("/:subject_attendance_id", handler.Get)
}
