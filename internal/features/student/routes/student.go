package routes

import (
	"api/internal/features/student/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterStudent(g *gin.RouterGroup, handler *handlers.Student) {
	newGroup := g.Group("/batches/:batch_id/majors/:major_id/classrooms/:classroom_id/students")
	newGroup.GET("", handler.GetAllByClassroomId)

	group := g.Group("/students")

	group.GET("", handler.GetAll)
}
