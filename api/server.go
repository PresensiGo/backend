package api

import (
	_ "api/docs"
	"api/internal/routes"
	"api/pkg/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func NewServer() {
	router := gin.Default()
	v1 := router.Group("/api/v1")

	routes.RegisterAuth(v1)

	// protected routes
	authorized := v1.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		routes.RegisterBatch(authorized)
		routes.RegisterClass(authorized)
		routes.RegisterMajor(authorized)
		routes.RegisterExcel(authorized)
		routes.RegisterStudent(authorized)
	}

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8080")
}
