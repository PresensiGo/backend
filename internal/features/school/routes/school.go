package routes

import (
	"api/internal/features/school/handlers"
	"api/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterSchool(g *gin.RouterGroup, handler *handlers.School) {
	group := g.Group("/schools").Use(middlewares.AuthMiddleware())

	group.GET("/profile", handler.GetSchool)
	group.PUT("/profile", handler.UpdateSchool)
}
