package main

import (
	"CLI-Tools/database"
	"CLI-Tools/routes"
	"net/http"
)

type userInterface struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {

	database.ConnectDatabase()
	//conn := database.GetDatabase()
	database.Migrate()
	mux := http.NewServeMux()

	FileUploadHandler := routes.HttpHandler{}

	mux.Handle("/getFileInformation", &FileUploadHandler)
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		panic(err)
	}
}
