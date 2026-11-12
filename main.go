
package main

import (
	"log"

	"backend-api/config"
	"backend-api/models"

	"github.com/joho/godotenv"
)

func main() {
	// Membaca konfigurasi .env
	err := godotenv.Load()
	if err != nil {
		log.Println("File .env tidak ditemukan")
	}

	// Menghubungkan database
	config.ConnectDatabase()

	// Membuat tabel secara otomatis
	err = config.DB.AutoMigrate(&models.Applicant{})
	if err != nil {
		log.Fatal("Gagal melakukan migrasi:", err)
	}

	log.Println("Tabel applicants berhasil disiapkan!")
}