package routes

import (
	"api/internal/features/subject/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterSubject(g *gin.RouterGroup, handler *handlers.Subject) {
	group := g.Group("/subjects")

	group.POST("", handler.Create)
	group.GET("", handler.GetAll)
	group.PUT("/:subject_id", handler.Update)
	group.DELETE("/:subject_id", handler.Delete)
}
