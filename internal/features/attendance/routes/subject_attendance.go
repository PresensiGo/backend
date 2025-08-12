package routes

import (
	"api/internal/features/attendance/handlers"
	"api/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterSubjectAttendance(g *gin.RouterGroup, handler *handlers.SubjectAttendance) {
	relativePath := "/batches/:batch_id/majors/:major_id/classrooms/:classroom_id/subject-attendances"
	{
		group := g.Group(relativePath).Use(middlewares.AuthMiddleware())
		group.POST("", handler.Create)
		group.GET("", handler.GetAll)
		group.GET("/:subject_attendance_id", handler.Get)
	}

	{
		group := g.Group("/subject-attendances").Use(middlewares.StudentMiddleware())
		group.POST("/records/student", handler.CreateRecordStudent)
	}
}
