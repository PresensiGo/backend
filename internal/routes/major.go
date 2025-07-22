package routes

import (
	"api/internal/injectors"
	"github.com/gin-gonic/gin"
)

func RegisterMajor(g *gin.RouterGroup) {
	group := g.Group("/major")
	handler := injectors.InitMajorHandler()

	group.GET("/batch/:batch_id", handler.GetAllMajors)
}
