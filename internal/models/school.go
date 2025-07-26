package models

import "gorm.io/gorm"

type School struct {
	gorm.Model

	Code string
	Name string
}
