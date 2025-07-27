package routes

import (
	"api/internal/injectors"
	"github.com/gin-gonic/gin"
)

func RegisterExcel(g *gin.RouterGroup) {
	group := g.Group("/excel")
	handler := injectors.InitExcelHandler()

	group.POST("/import", handler.ImportData)
	group.POST("/import-v2", handler.ImportDataV2)
}
