package models

import (
	models2 "api/internal/features/school/models"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string
	Email    string `gorm:"unique"`
	Password string
	Role     string `gorm:"default:'student'"`
	SchoolId uint
	School   models2.School
}
