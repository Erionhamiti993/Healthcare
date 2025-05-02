package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	PatientID uint      `gorm:"not null" validate:"required"`
	DateTime  time.Time `gorm:"not null" validate:"required"`
	Reason    string    `gorm:"not null" validate:"required"`
	Patient   Patient   `gorm:"foreignKey:PatientID"`
}
