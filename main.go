package main

import (
	"CLI-Tools/database"
	"CLI-Tools/routes"
	"net/http"
)

func main() {

	database.ConnectDatabase()
	conn := database.GetDatabase()

	_ = conn
	// database.Migrate()
	database.Migrate()

	mux := http.NewServeMux()

	FileUploadHandler := routes.HttpHandler{}
	AuthHandler := routes.AuthenticationHandler{}
	LogHandler := routes.MediaHandler{}

	mux.Handle("/auth/", &AuthHandler)
	mux.Handle("/upload", &FileUploadHandler)
	mux.HandleFunc("/video/", LogHandler.ServeHTTP)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
