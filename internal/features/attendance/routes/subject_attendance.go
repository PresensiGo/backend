package routes

import (
	"api/internal/features/attendance/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterSubjectAttendance(g *gin.RouterGroup, handler *handlers.SubjectAttendance) {
	newGroup := g.Group("/batches/:batch_id/majors/:major_id/classrooms/:classroom_id/subject-attendances")
	newGroup.POST("", handler.Create)
	newGroup.GET("", handler.GetAll)
	newGroup.GET("/:subject_attendance_id", handler.Get)
}
