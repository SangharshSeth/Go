package main

import (
	"github.com/sangharshseth/routes"
	"log"
	"net/http"
	"os"
)

func main() {

	mux := http.NewServeMux()

	FileUploadHandler := routes.HttpHandler{}
	AuthHandler := routes.AuthenticationHandler{}
	LogHandler := routes.MediaHandler{}

	//Routes
	mux.Handle("/auth/", &AuthHandler)
	mux.Handle("/upload", &FileUploadHandler)
	mux.HandleFunc("/video/", LogHandler.ServeHTTP)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error %s", err.Error())
		os.Exit(0)
	}
}
