package classroom

import (
	"api/internal/features/classroom/routes"
	"api/internal/injector"
	"github.com/gin-gonic/gin"
)

func RegisterClassroom(g *gin.RouterGroup) {
	handlers := injector.InitClassroomHandlers()

	routes.RegisterClassroom(g, handlers.Classroom)
}
