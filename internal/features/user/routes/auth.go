package routes

import (
	"api/internal/features/user/handlers"
	"api/pkg/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterAuth(g *gin.RouterGroup, handler *handlers.Auth) {
	{
		group := g.Group("/auth")

		group.POST("/login", handler.Login)
		group.POST("/refresh-token", handler.RefreshToken)
	}

	{
		group := g.Group("/auth").Use(middlewares.AuthMiddleware())

		group.POST("/logout", handler.Logout)
	}
}
