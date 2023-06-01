package database

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	db *gorm.DB
)

func ConnectDatabase() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Failed to Load Environment Variables")
	}
	DatabaseConnectionString := os.Getenv("DSN")
	conn, err := gorm.Open(postgres.Open(DatabaseConnectionString),
		&gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = conn
}

func GetDatabase() *gorm.DB {
	return db
}
