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

	routes.RegisterAuthRoutes(v1)

	// protected routes
	authorized := v1.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		routes.RegisterBatchRoutes(authorized)
		routes.RegisterClassRoutes(authorized)
		routes.RegisterMajorRoutes(authorized)
		routes.RegisterExcelRoutes(authorized)
	}

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8080")
}
