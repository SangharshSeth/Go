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

	mux := http.NewServeMux()

	FileUploadHandler := routes.HttpHandler{}
	AuthHandler := routes.AuthenticationHandler{}
	mux.Handle("/signup", &AuthHandler)
	mux.Handle("/upload", &FileUploadHandler)

	err := http.ListenAndServe(":80", mux)
	if err != nil {
		panic(err)
	}
}
