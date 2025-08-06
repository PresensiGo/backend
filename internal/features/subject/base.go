package subject

import (
	"api/internal/features/subject/routes"
	"api/internal/injector"
	"github.com/gin-gonic/gin"
)

func RegisterSubject(g *gin.RouterGroup) {
	handlers := injector.InitSubjectHandlers()

	routes.RegisterSubject(g, handlers.Subject)
}
