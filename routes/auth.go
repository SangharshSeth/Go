package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sangharshseth/database"
	"github.com/sangharshseth/models"
)

type AuthenticationHandler struct{}

type UserDetails struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(writer http.ResponseWriter, request *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Print("Failed to Parse the env guys")
		return
	}
	database.ConnectDatabase()
	ctx, db := database.GetDatabase()
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	var data UserDetails
	bodyParseError := decoder.Decode(&data)
	if bodyParseError != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	user := models.User{
		Email:    data.Email,
		Password: data.Password,
	}
	coll := db.Database("development").Collection("Users")
	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		log.Printf("Failed to insert %s", err.Error())
		_, err2 := writer.Write([]byte("Failed to insert"))
		if err2 != nil {
			return
		}
		os.Exit(0)
	}

	respBody := make(map[string]string)
	respBody["data"] = "User Successfully Created"
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Token", "Sangharsh")
	err = json.NewEncoder(writer).Encode(&respBody)

	if err != nil {
		// Handle the error
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (auth *AuthenticationHandler) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	log.Print(request.URL.Path)
	switch request.URL.Path {
	case "/auth/signup":
		Signup(write, request)
	}

}
