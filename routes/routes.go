package routes

import (
	"healthcare/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Logging middleware
	router.Use(gin.Logger())

	// Patient routes
	patientRoutes := router.Group("/patients")
	{
		patientRoutes.POST("", handlers.CreatePatient)
		patientRoutes.GET("", handlers.GetPatients)
		patientRoutes.GET("/:id", handlers.GetPatientByID)
	}

	// Appointment routes
	appointmentRoutes := router.Group("/appointments")
	{
		appointmentRoutes.POST("", handlers.CreateAppointment)
		appointmentRoutes.GET("", handlers.GetAppointments)
	}

	return router
}
