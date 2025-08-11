package teacher

import (
	"api/internal/features/teacher/routes"
	"github.com/gin-gonic/gin"
)

func RegisterModule(g *gin.RouterGroup) {
	routes.RegisterAuth(g, nil)
}
