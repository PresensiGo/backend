package routes

import (
	"api/internal/features/batch/handlers"
	"api/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterBatch(g *gin.RouterGroup, handler *handlers.Batch) {
	group := g.Group("/batches").Use(middlewares.AuthMiddleware())

	group.POST("", handler.Create)
	group.GET("", handler.GetAllBatches)
	group.GET("/:batch_id", handler.Get)
	group.PUT("/:batch_id", handler.Update)
	group.DELETE("/:batch_id", handler.Delete)
}
