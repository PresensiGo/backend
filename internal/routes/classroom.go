package routes

import (
	"api/internal/injectors"
	"github.com/gin-gonic/gin"
)

func RegisterClassroom(g *gin.RouterGroup) {
	group := g.Group("/classrooms")
	handler := injectors.InitClassroomHandler()

	group.GET("/batches/:batch_id", handler.GetAllWithMajors)
}
