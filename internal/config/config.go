package config

import (
	"fmt"
	"github.com/MaLuMSiza/go-crud-gin/internal/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("falha ao carregar .env: %w", err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar ao banco: %w", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, fmt.Errorf("falha ao migrar banco de dados: %w", err)
	}

	return db, nil
}
