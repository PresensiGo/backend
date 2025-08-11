package user

import (
	"fmt"

	"api/internal/features/user/routes"
	"api/internal/injector"
	"github.com/gin-gonic/gin"
)

func RegisterUser(g *gin.RouterGroup) {
	handlers := injector.InitUserHandlers()

	routes.RegisterUser(g, handlers.User)
	routes.RegisterAuth(g, handlers.Auth)

	// helpers
	if err := handlers.Admin.Inject(); err != nil {
		fmt.Println(err.Error())
	}
}
