package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectDatabase() {
	conn, err := gorm.Open(postgres.Open("postgresql://sangharsh:JsywqQBG_W8fQw7P-ZEFAg@development-2927.7s5.cockroachlabs.cloud:26257/development?sslmode=verify-full"),
		&gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = conn
}

func GetDatabase() *gorm.DB {
	return db
}
