package routes

import (
	"api/internal/features/attendance/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterLateness(g *gin.RouterGroup, handler *handlers.Lateness) {
	group := g.Group("/latenesses")

	group.POST("", handler.Create)
	group.GET("", handler.GetAll)

	group.POST("/:lateness_id", handler.CreateDetail)
	group.GET("/:lateness_id", handler.Get)
}
