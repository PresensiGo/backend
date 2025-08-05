package routes

import (
	"api/internal/features/classroom/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterClassroom(g *gin.RouterGroup, handler *handlers.Classroom) {
	group := g.Group("/classrooms")

	group.GET("", handler.GetAll)
	group.GET("/batches/:batch_id", handler.GetAllWithMajors)
}
