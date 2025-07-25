package routes

import (
	"api/internal/injectors"
	"github.com/gin-gonic/gin"
)

func RegisterStudent(g *gin.RouterGroup) {
	group := g.Group("/students")
	handler := injectors.InitStudentHandler()

	group.GET("/classrooms/:classroom_id", handler.GetAll)
}
