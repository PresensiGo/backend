package routes

import (
	"api/internal/features/user/handlers"
	"api/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterUser(g *gin.RouterGroup, handler *handlers.User) {
	group := g.Group("/accounts").Use(middlewares.AuthMiddleware())

	group.POST("/import", handler.ImportAccounts)
	group.GET("", handler.GetAll)
	group.PUT("/:account_id/password", handler.UpdateAccountPassword)
	group.DELETE("/:account_id", handler.DeleteAccount)
}
