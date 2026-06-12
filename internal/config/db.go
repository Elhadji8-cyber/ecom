package config

import (
	"fmt"
	"log"
	"os"

	authModels "server/internal/auth/models"
	productModels "server/internal/product/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = database.AutoMigrate(
		&authModels.Customer{},
		&productModels.Category{},
		&productModels.Product{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = database
	fmt.Println("Database connection established and migrated")
}
