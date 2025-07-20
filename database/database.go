package database

import (
	models2 "api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type databaseConfig struct {
	DB *gorm.DB
}

func NewDatabaseConfig() *databaseConfig {
	dsn := "host=localhost user=root dbname=presensi_sekolah port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	db.AutoMigrate(
		&models2.User{},
		&models2.UserToken{},
		&models2.Batch{},
		&models2.Major{},
		&models2.Class{},
		&models2.Student{},
	)

	return &databaseConfig{
		DB: db,
	}
}
