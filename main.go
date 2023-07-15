package main

import (
	"fmt"
	"github.com/sangharshseth/routes"
	"log"
	"net/http"
	"os"
)

func main() {

	//conn := database.GetDatabase()

	//database.Migrate()
	//database.Migrate()

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
	})
	mux.Handle("/auth/", &AuthHandler)
	mux.Handle("/upload", &FileUploadHandler)
	mux.HandleFunc("/video/", LogHandler.ServeHTTP)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error %s", err.Error())
		os.Exit(0)
	}
}
