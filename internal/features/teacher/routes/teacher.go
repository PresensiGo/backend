package routes

import (
	"api/internal/features/teacher/handlers"
	"api/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterTeacher(g *gin.RouterGroup, handler *handlers.Teacher) {
	group := g.Group("/teachers").Use(middlewares.AuthMiddleware())
	group.POST("/import", handler.Import)
}
