package batch

import (
	"api/internal/features/batch/routes"
	"api/internal/injector"
	"github.com/gin-gonic/gin"
)

func RegisterBatch(g *gin.RouterGroup) {
	handlers := injector.InitBatchHandlers()

	routes.RegisterBatch(g, handlers.Batch)
}
