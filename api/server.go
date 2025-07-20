package api

import (
	"api/database"
	_ "api/docs"
	"api/features/batch"
	"api/features/excel"
	"api/handler"
	"api/internal/handlers"
	"api/internal/routes"
	"api/internal/services"
	"api/pkg/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func NewServer() {
	router := gin.Default()
	db := database.NewDatabaseConfig()

	v1 := router.Group("/api/v1")

	routes.RegisterAuthRoutes(v1, db.DB)

	// protected routes
	authorized := v1.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		// excel
		{
			excelRouter := authorized.Group("/excel")
			excelService := excel.NewService(db.DB)
			excelHandler := handler.NewExcelHandler(excelService)

			excelRouter.POST("/import", excelHandler.Import)
		}

		// batch
		{
			batchRouter := authorized.Group("/batch")
			batchService := batch.NewService(db.DB)
			batchHandler := handler.NewBatchHandler(batchService)

			batchRouter.POST("/", batchHandler.CreateBatch)
			//batchRouter.GET("/", batchHandler.GetAllBatches)
			//batchRouter.GET("/:id", batchHandler.GetBatch)
			//batchRouter.PUT("/:id", batchHandler.UpdateBatch)
			//batchRouter.DELETE("/:id", batchHandler.DeleteBatch)
		}

		// major

		// class

		// reset
		{
			resetRouter := authorized.Group("/reset")
			resetService := services.NewResetService(db.DB)
			resetHandler := handlers.NewResetHandler(resetService)

			resetRouter.POST("/", resetHandler.Reset)
		}
	}

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8080")
}
