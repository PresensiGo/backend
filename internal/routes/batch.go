package routes

import (
	"api/internal/injectors"
	"github.com/gin-gonic/gin"
)

func RegisterBatch(g *gin.RouterGroup) {
	group := g.Group("/batch")
	handler := injectors.InitBatchHandler()

	group.GET("/", handler.GetAll)
}
