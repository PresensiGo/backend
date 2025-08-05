package routes

import (
	"api/internal/features/data/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterReset(g *gin.RouterGroup, handler *handlers.Reset) {
	group := g.Group("/reset")

	group.GET("", handler.Reset)
}
