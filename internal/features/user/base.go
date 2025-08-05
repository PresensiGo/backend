package user

import (
	"api/internal/features/user/routes"
	"api/internal/injector"
	"github.com/gin-gonic/gin"
)

func RegisterUser(g *gin.RouterGroup) {
	handlers := injector.InitUserHandlers()

	routes.RegisterAuth(g, handlers.Auth)
}
