package routes

import (
	"api/internal/injectors"
	"github.com/gin-gonic/gin"
)

func RegisterLateness(g *gin.RouterGroup) {
	group := g.Group("/latenesses")
	handler := injectors.InitLatenessHandler()

	group.POST("", handler.Create)
	group.GET("", handler.GetAll)
}
