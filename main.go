package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/sangharshseth/internal/connections"
	"github.com/sangharshseth/internal/routes"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Failed to Load Environment Variables")
	}

	MongoClient, err := connections.ConnectDatabase()
	if err != nil {
		log.Fatalf("Database Connection Failed: %s", err)
	}

	mux := http.NewServeMux()
	corsHandler := cors.Default().Handler(mux)

	FileUploadHandler := routes.HttpHandler{}
	AuthHandler := routes.AuthenticationHandler{
		Ctx: MongoClient.Ctx,
		Db:  MongoClient.Client,
	}

	//Routes
	mux.Handle("POST /auth/", &AuthHandler)
	mux.Handle("POST /upload", &FileUploadHandler)

	mux.HandleFunc("GET /healthcheck", healthCheck)

	err = http.ListenAndServe(":8080", corsHandler)
	if err != nil {
		log.Printf("Error %s", err.Error())
		return
	}
}
