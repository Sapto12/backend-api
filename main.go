
package main

import (
	"log"
	"os"

	"backend-api/config"
	"backend-api/models"
	"backend-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("File .env tidak ditemukan, menggunakan environment sistem")
	}

	// Connect database
	config.ConnectDatabase()

	// Database migration
	if err := config.DB.AutoMigrate(&models.Applicant{}); err != nil {
		log.Fatal("Gagal melakukan migrasi database:", err)
	}

	log.Println("Database migration berhasil")

	// Initialize Gin
	router := gin.Default()

	// Register routes
	routes.SetupRoutes(router)

	// Get application port
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Start server
	log.Println("Server berjalan di port:", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}
}
