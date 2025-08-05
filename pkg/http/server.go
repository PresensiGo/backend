package http

import (
	_ "api/docs"
	"api/internal/features/attendance"
	"api/internal/features/batch"
	"api/internal/features/classroom"
	"api/internal/features/data"
	"api/internal/features/major"
	"api/internal/features/student"
	"api/internal/features/user"
	"api/pkg/http/middleware"
	"github.com/gin-contrib/cors"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func NewServer() {
	router := gin.Default()

	// cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	corsConfig.AllowCredentials = true

	router.Use(cors.New(corsConfig))

	v1 := router.Group("/api/v1")

	user.RegisterUser(v1)

	// protected routes
	authorized := v1.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		attendance.RegisterAttendance(authorized)
		batch.RegisterBatch(authorized)
		major.RegisterMajor(authorized)
		classroom.RegisterClassroom(authorized)
		student.RegisterStudent(authorized)
		data.RegisterData(authorized)
	}

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8080")
}
