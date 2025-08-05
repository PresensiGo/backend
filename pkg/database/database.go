package database

import (
	"fmt"
	"os"
	"sync"

	models3 "api/internal/features/attendance/models"
	models4 "api/internal/features/batch/models"
	models6 "api/internal/features/classroom/models"
	models5 "api/internal/features/major/models"
	models2 "api/internal/features/school/models"
	models7 "api/internal/features/student/models"
	models8 "api/internal/features/user/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

func New() *gorm.DB {
	dbOnce.Do(
		func() {
			dbHost := os.Getenv("DB_HOST")
			dbPort := os.Getenv("DB_PORT")
			dbName := os.Getenv("DB_NAME")
			dbUser := os.Getenv("DB_USER")
			dbPassword := os.Getenv("DB_PASSWORD")
			dbTimezone := os.Getenv("DB_TIMEZONE")

			dsn := fmt.Sprintf(
				"host=%s user=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
				dbHost, dbUser, dbName, dbPort, dbTimezone,
			)
			if dbPassword != "" {
				dsn = fmt.Sprintf(
					"%s password=%s",
					dsn, dbPassword,
				)
			}

			db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err != nil {
				panic("failed to connect to database: " + err.Error())
			}

			_ = db.AutoMigrate(
				&models8.User{},
				&models8.UserToken{},
				&models2.School{},
				&models4.Batch{},
				&models5.Major{},
				&models6.Classroom{},
				&models7.Student{},
				&models3.Attendance{},
				&models3.AttendanceDetail{},
				&models3.Lateness{},
				&models3.LatenessDetail{},
			)

			dbInstance = db
		},
	)

	return dbInstance
}
