package routes

import (
	"api/internal/features/major/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterMajor(g *gin.RouterGroup, handler *handlers.Major) {
	newGroup := g.Group("/batches/:batch_id/majors")
	newGroup.GET("", handler.GetAllByBatchId)

	group := g.Group("/majors")

	group.POST("", handler.Create)
	group.GET("", handler.GetAllMajors)
	group.PUT("/:major_id", handler.Update)
	group.DELETE("/:major_id", handler.Delete)
}
