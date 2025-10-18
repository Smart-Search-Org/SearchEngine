package utils

import (
	"SmartSearch/internal/models/database"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&database.UserIndex{})
	if err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	DB = db
	log.Println("Database connected and migrated successfully!")
}
