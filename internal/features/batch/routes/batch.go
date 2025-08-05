package routes

import (
	"api/internal/features/batch/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterBatch(g *gin.RouterGroup, handler *handlers.Batch) {
	group := g.Group("/batches")

	group.POST("", handler.Create)
	group.GET("", handler.GetAll)
	group.PUT("/:batch_id", handler.Update)
	group.DELETE("/:batch_id", handler.Delete)
}
