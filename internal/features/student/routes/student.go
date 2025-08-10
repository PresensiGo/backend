package routes

import (
	"api/internal/features/student/handlers"
	"api/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterStudent(g *gin.RouterGroup, handler *handlers.Student) {
	newGroup := g.Group("/batches/:batch_id/majors/:major_id/classrooms/:classroom_id/students").Use(middlewares.AuthMiddleware())
	newGroup.GET("", handler.GetAllByClassroomId)

	// old
	group := g.Group("/students")
	group.GET("", handler.GetAll)
}
