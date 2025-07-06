package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"social-media-service/internal/model"
)

// DB is the global database connection
var DB *gorm.DB

// Init connects to the database, migrates schema, and seeds data.
func Connect() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to DB: %w", err)
	}

	// Auto-migrate the User model
	if err := DB.AutoMigrate(&model.User{}); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	// Seed data if no users exist
	var count int64
	if err := DB.Model(&model.User{}).Count(&count).Error; err != nil {
		return fmt.Errorf("count failed: %w", err)
	}

	if count == 0 {
		users := []model.User{
			{Username: "alice", Email: "alice@example.com", Password: "hashedpassword1"},
			{Username: "bob", Email: "bob@example.com", Password: "hashedpassword2"},
		}
		if err := DB.Create(&users).Error; err != nil {
			return fmt.Errorf("failed to seed users: %w", err)
		}
		log.Println("Seeded initial users.")
	} else {
		log.Println("Users already exist. Skipping seeding.")
	}

	log.Println("Database initialized successfully.")
	return nil
}
