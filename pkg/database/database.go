package database

import (
	"fmt"
	"os"
	"sync"

	attendance "api/internal/features/attendance/models"
	batch "api/internal/features/batch/models"
	classroom "api/internal/features/classroom/models"
	major "api/internal/features/major/models"
	school "api/internal/features/school/models"
	student "api/internal/features/student/models"
	subject "api/internal/features/subject/models"
	user "api/internal/features/user/models"
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
				&user.User{},
				&user.UserToken{},
				&user.UserSession{},
				&school.School{},
				&batch.Batch{},
				&major.Major{},
				&classroom.Classroom{},
				&student.Student{},

				&attendance.GeneralAttendance{},
				&attendance.GeneralAttendanceRecord{},
				&attendance.SubjectAttendance{},
				&attendance.SubjectAttendanceRecord{},
				&subject.Subject{},
				&student.StudentToken{},
			)

			dbInstance = db
		},
	)

	return dbInstance
}
