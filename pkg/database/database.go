package database

import (
	"api/internal/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func New() *gorm.DB {
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

	fmt.Println("dsn", dsn)

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
