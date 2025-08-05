package routes

import (
	"api/internal/features/student/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterStudent(g *gin.RouterGroup, handler *handlers.Student) {
	group := g.Group("/students")

	group.GET("", handler.GetAll)
	group.GET("/classrooms/:classroom_id", handler.GetAllByClassroomId)
}
