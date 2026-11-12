
package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		log.Fatal("DATABASE_URL belum dikonfigurasi")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal menghubungkan database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Gagal mendapatkan database connection:", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Database tidak dapat diakses:", err)
	}

	DB = db

	fmt.Println("Database Neon PostgreSQL berhasil terhubung!")
}