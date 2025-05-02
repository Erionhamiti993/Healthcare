package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	Name         string        `gorm:"not null" validate:"required"`
	Email        string        `gorm:"unique;not null" validate:"required,email"`
	Appointments []Appointment `gorm:"foreignKey:PatientID"`
}
