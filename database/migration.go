package database

import "CLI-Tools/models"

func Migrate() {
	ConnectDatabase()
	db := GetDatabase()
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
}
