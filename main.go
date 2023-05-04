package main

import (
	"CLI-Tools/database"
	"CLI-Tools/routes"
	"fmt"
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

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fprint, err := fmt.Fprint(writer, "Hello From Port 8080")
		if err != nil {
			return
		}
		_ = fprint
		return
	})
	mux.Handle("/auth/", &AuthHandler)
	mux.Handle("/upload", &FileUploadHandler)
	mux.HandleFunc("/video/", LogHandler.ServeHTTP)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
