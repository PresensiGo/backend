package routes

import (
	"api/internal/features/data/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterExcel(g *gin.RouterGroup, handler *handlers.Excel) {
	group := g.Group("/excel")

	// group.POST("/import", handler.ImportData)
	// group.POST("/import-v2", handler.ImportDataV2)
	group.POST("/import-data", handler.ImportDataV3)
}
