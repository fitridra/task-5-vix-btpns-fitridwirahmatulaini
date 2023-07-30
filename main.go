package main

import (
	"btpn-fitri/database"
	"btpn-fitri/models"
	"btpn-fitri/router"
)

func main() {
	db := database.SetupDB()
	db.AutoMigrate(&models.User{})

	r := router.SetupRoutes(db)
	r.Run()
}
