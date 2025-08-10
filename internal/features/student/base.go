package student

import (
	"api/internal/features/student/routes"
	"api/internal/injector"
	"github.com/gin-gonic/gin"
)

func RegisterStudent(g *gin.RouterGroup) {
	handlers := injector.InitStudentHandlers()

	routes.RegisterStudent(g, handlers.Student)
	routes.RegisterStudentAuth(g, handlers.StudentAuth)
}
