package routes

import (
	"api/internal/injectors"
	"github.com/gin-gonic/gin"
)

func RegisterStudent(g *gin.RouterGroup) {
	group := g.Group("/student")
	handler := injectors.InitStudentHandler()

	group.GET("/class/:class_id", handler.GetAll)
}
