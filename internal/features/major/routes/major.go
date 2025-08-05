package routes

import (
	"api/internal/features/major/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterMajor(g *gin.RouterGroup, handler *handlers.Major) {
	group := g.Group("/majors")

	group.GET("/batch/:batch_id", handler.GetAll)
}
