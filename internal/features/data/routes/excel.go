package routes

import (
	"api/internal/features/data/handlers"
	"api/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterExcel(g *gin.RouterGroup, handler *handlers.Excel) {
	group := g.Group("/excel").Use(middlewares.AuthMiddleware())

	group.POST("/import-data", handler.ImportData)
}
