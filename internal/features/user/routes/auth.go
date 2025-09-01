package routes

import (
	"api/internal/features/user/handlers"
	"api/pkg/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterAuth(g *gin.RouterGroup, handler *handlers.Auth) {
	relativePath := "/auth"
	{
		group := g.Group(relativePath)

		group.POST("/login", handler.Login)
		group.POST("/login-2", handler.Login2)
		group.POST("/refresh-token", handler.RefreshToken)
	}
	{
		group := g.Group(relativePath).Use(middlewares.AuthMiddleware())

		group.POST("/logout", handler.Logout)
		group.POST("/logout-2", handler.Logout2)
	}
}
