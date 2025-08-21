package school

import (
	"api/internal/features/school/routes"
	"api/internal/injector"
	"github.com/gin-gonic/gin"
)

func RegisterModule(g *gin.RouterGroup) {
	handlers := injector.InitSchoolHandlers()

	routes.RegisterSchool(g, handlers.School)
}
