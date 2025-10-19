package utils

import (
	"SmartSearch/internal/models/database"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&database.UserIndex{})
	if err != nil {
		log.Printf("Database migration failed (likely that the table already exists): %v\n", err)
	}

	DB = db
	log.Println("Database connected and migrated successfully!")
}
