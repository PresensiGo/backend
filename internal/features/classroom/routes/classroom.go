package routes

import (
	"api/internal/features/classroom/handlers"
	"api/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterClassroom(g *gin.RouterGroup, handler *handlers.Classroom) {
	{
		relativePath := "/batches/:batch_id/majors/:major_id/classrooms"
		group := g.Group(relativePath).Use(middlewares.AuthMiddleware())

		group.POST("", handler.Create)
		group.GET("", handler.GetAllClassroomsByMajorId)
		group.PUT("/:classroom_id", handler.Update)
	}
	{
		group := g.Group("/classrooms").Use(middlewares.AuthMiddleware())

		group.GET("", handler.GetAll)
		group.GET("/batches/:batch_id", handler.GetAllWithMajors)
	}
}
