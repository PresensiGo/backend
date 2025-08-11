package teacher

import (
	"api/internal/features/teacher/routes"
	"api/internal/injector"
	"github.com/gin-gonic/gin"
)

func RegisterModule(g *gin.RouterGroup) {
	handlers := injector.InitTeacherHandlers()

	routes.RegisterTeacher(g, handlers.Teacher)
}
