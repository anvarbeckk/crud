package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres dbname=crud port=5432 sslmode=disable timezone=Asia/Tashkent"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}
	// register post model
	database.AutoMigrate(&Post{})

	DB = database
}