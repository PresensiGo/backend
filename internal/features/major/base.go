package major

import (
	"api/internal/features/major/routes"
	"api/internal/injector"
	"github.com/gin-gonic/gin"
)

func RegisterMajor(g *gin.RouterGroup) {
	handlers := injector.InitMajorHandlers()

	routes.RegisterMajor(g, handlers.Major)
}
