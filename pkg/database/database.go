package database

import (
	"api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	dsn := "host=localhost user=root dbname=presensi_sekolah port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	db.AutoMigrate(
		&models.User{},
		&models.UserToken{},
		&models.Batch{},
		&models.Major{},
		&models.Class{},
		&models.Student{},
	)

	return db
}
