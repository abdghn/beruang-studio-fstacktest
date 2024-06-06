package models

import "gorm.io/gorm"

type Invitation struct {
	gorm.Model
	Email  string
	Link   string
	Status string
}
