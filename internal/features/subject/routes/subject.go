package routes

import (
	"api/internal/features/subject/handlers"
	"api/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterSubject(g *gin.RouterGroup, handler *handlers.Subject) {
	group := g.Group("/subjects").Use(middlewares.AuthMiddleware())

	group.POST("", handler.Create)
	group.GET("", handler.GetAllSubjects)
	group.PUT("/:subject_id", handler.Update)
	group.DELETE("/:subject_id", handler.Delete)
}
