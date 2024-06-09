package db

import (
	"log"
	"os"

	"grpc-go-proto-product-ms-main/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init() Handler {
	url := os.Getenv("DB_URL")
	if url == "" {
		log.Fatalln("DB_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.StockDecreaseLog{})

	return Handler{DB: db}
}

