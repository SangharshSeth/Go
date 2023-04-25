package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectDatabase() {
	conn, err := gorm.Open(mysql.Open("root:sangharsh@tcp(127.0.0.1)/development?utf8mb4&parseTime=True&loc=Local&tls=false"),
		&gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = conn
}

func GetDatabase() *gorm.DB {
	return db
}
