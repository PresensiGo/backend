package routes

import (
	"api/internal/injectors"
	"github.com/gin-gonic/gin"
)

func RegisterReset(g *gin.RouterGroup) {
	group := g.Group("/reset")
	handler := injectors.InitResetHandler()

	group.GET("", handler.Reset)
}
