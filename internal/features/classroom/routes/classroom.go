package routes

import (
	"api/internal/features/classroom/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterClassroom(g *gin.RouterGroup, handler *handlers.Classroom) {
	newGroup := g.Group("/batches/:batch_id/majors/:major_id/classrooms")
	newGroup.POST("", handler.Create)
	newGroup.GET("", handler.GetAllByMajorId)

	group := g.Group("/classrooms")

	group.GET("", handler.GetAll)
	group.GET("/batches/:batch_id", handler.GetAllWithMajors)
}
