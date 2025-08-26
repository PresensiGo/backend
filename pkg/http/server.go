package http

import (
	_ "api/docs"
	"api/internal/features/attendance"
	"api/internal/features/batch"
	"api/internal/features/classroom"
	"api/internal/features/data"
	"api/internal/features/major"
	"api/internal/features/school"
	"api/internal/features/student"
	"api/internal/features/subject"
	"api/internal/features/user"
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

	batch.RegisterBatch(v1)
	major.RegisterMajor(v1)
	classroom.RegisterClassroom(v1)
	attendance.RegisterAttendance(v1)
	user.RegisterUser(v1)
	student.RegisterStudent(v1)
	subject.RegisterSubject(v1)
	school.RegisterModule(v1)
	data.RegisterModule(v1)

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8080")
}
