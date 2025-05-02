package main

import (
	"healthcare/database"
	"healthcare/routes"
)

func main() {
	database.InitDB()
	router := routes.SetupRouter()
	router.Run(":8080")
}
