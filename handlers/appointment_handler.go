package handlers

import (
	"net/http"

	"healthcare/database"
	"healthcare/models"

	"github.com/gin-gonic/gin"
)

func CreateAppointment(c *gin.Context) {
	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if patient exists
	var patient models.Patient
	result := database.DB.First(&patient, appointment.PatientID)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Patient not found"})
		return
	}

	result = database.DB.Create(&appointment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, appointment)
}

func GetAppointments(c *gin.Context) {
	var appointments []models.Appointment
	query := database.DB.Preload("Patient")

	if patientID := c.Query("patient_id"); patientID != "" {
		query = query.Where("patient_id = ?", patientID)
	}

	result := query.Find(&appointments)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, appointments)
}
