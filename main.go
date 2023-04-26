package main

import (
	"CLI-Tools/database"
	"CLI-Tools/models"
	"CLI-Tools/routes"
	"net/http"
)

func main() {

	database.ConnectDatabase()
	conn := database.GetDatabase()
	UserProfile := models.UserProfile{
		Sex:    "AlphaMale",
		Age:    20,
		Height: 5.11,
		Weight: 60,
		Likes:  []string{"girls", "videogames"},
	}
	database.Migrate()
	conn.Create(&UserProfile)

	mux := http.NewServeMux()

	FileUploadHandler := routes.HttpHandler{}

	mux.Handle("/getFileInformation", &FileUploadHandler)
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		panic(err)
	}
}
