package main

import (
	"log"

    "github.com/joho/godotenv"
	"myapp/config"
	"myapp/models"
	"myapp/routes"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Println("Warning: .env file not found, relying on system environment variables")
    }
	
	// Koneksi DB
	config.ConnectDatabase()

	// Auto migrate tabel
	config.DB.AutoMigrate(&models.User{})

	// Setup routes
	r := routes.SetupRouter()
	r.Run(":8080") // jalan di localhost:8080
}
