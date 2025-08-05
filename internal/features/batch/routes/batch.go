package routes

import (
	"api/internal/features/batch/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterBatch(g *gin.RouterGroup, handler *handlers.Batch) {
	group := g.Group("/batch")

	group.GET("", handler.GetAll)
}
