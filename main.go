package main

import (
	"github.com/sangharshseth/database"
	"github.com/sangharshseth/routes"
	"log"
	"net/http"
	"os"
)

func main() {

	database.ConnectDatabase()
	ctx, db := database.GetDatabase()
	mux := http.NewServeMux()

	FileUploadHandler := routes.HttpHandler{}
	AuthHandler := routes.AuthenticationHandler{
		Ctx: ctx,
		Db:  db,
	}

	//Routes
	mux.Handle("/auth/", &AuthHandler)
	mux.Handle("/upload", &FileUploadHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error %s", err.Error())
		os.Exit(0)
	}
}
