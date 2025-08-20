package routes

import (
	"api/internal/features/major/handlers"
	"api/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterMajor(g *gin.RouterGroup, handler *handlers.Major) {
	{
		relativePath := "/batches/:batch_id/majors"
		group := g.Group(relativePath).Use(middlewares.AuthMiddleware())

		group.POST("", handler.Create)
		group.GET("", handler.GetAllMajorsByBatchId)
		group.GET("/:major_id", handler.GetMajor)
		group.PUT("/:major_id", handler.Update)
		group.DELETE("/:major_id", handler.Delete)
	}

	{
		group := g.Group("/majors").Use(middlewares.AuthMiddleware())

		group.GET("", handler.GetAllMajors)
	}
}
