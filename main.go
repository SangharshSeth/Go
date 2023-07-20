package main

import (
	"github.com/rs/cors"
	"github.com/sangharshseth/internal/connections"
	"github.com/sangharshseth/internal/routes"
	"log"
	"net/http"
)

func main() {

	connections.ConnectDatabase()
	defer connections.CloseDatabase()

	ctx, db := connections.GetDatabase()

	mux := http.NewServeMux()

	corsHandler := cors.Default().Handler(mux)

	FileUploadHandler := routes.HttpHandler{}
	AuthHandler := routes.AuthenticationHandler{
		Ctx: ctx,
		Db:  db,
	}

	//Routes
	mux.Handle("/auth/", &AuthHandler)
	mux.Handle("/upload", &FileUploadHandler)

	err := http.ListenAndServe(":8080", corsHandler)
	if err != nil {
		log.Printf("Error %s", err.Error())
		return
	}
}
