package routes

import (
	"api/internal/injectors"
	"github.com/gin-gonic/gin"
)

func RegisterClassroom(g *gin.RouterGroup) {
	group := g.Group("/class")
	handler := injectors.InitClassHandler()

	group.GET("/major/:major_id", handler.GetAll)
}
