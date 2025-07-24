package routes

import (
	"api/internal/injectors"
	"github.com/gin-gonic/gin"
)

func RegisterClassMajor(g *gin.RouterGroup) {
	group := g.Group("/class_majors")
	handler := injectors.InitClassMajorHandler()

	group.GET("/batch/:batch_id", handler.GetAll)
}
