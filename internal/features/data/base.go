package data

import (
	"api/internal/features/data/routes"
	"api/internal/injector"
	"github.com/gin-gonic/gin"
)

func RegisterModule(g *gin.RouterGroup) {
	handlers := injector.InitDataHandlers()

	routes.RegisterExcel(g, handlers.Excel)
	// routes.RegisterReset(g, handlers.Reset)
}
