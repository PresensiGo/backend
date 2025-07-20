package models

import "gorm.io/gorm"

type Batch struct {
	gorm.Model

	Name string `json:"name" gorm:"not null;default:'Unnamed Batch'"`
}
