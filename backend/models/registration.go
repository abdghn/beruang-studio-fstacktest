package models

import "gorm.io/gorm"

type Registration struct {
	gorm.Model
	Email            string
	RegistrationCode string
	QRCode           string
	Status           string
}
