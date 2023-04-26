package database

import "CLI-Tools/models"

func Migrate() {
	ConnectDatabase()
	db := GetDatabase()
	err := db.AutoMigrate(&models.User{}, &models.UserProfile{})
	if err != nil {
		panic(err)
	}
}
