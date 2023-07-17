package routes

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sangharshseth/models"
)

type AuthenticationHandler struct {
	Ctx context.Context
	Db  *mongo.Client
}

type UserDetails struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(writer http.ResponseWriter, request *http.Request, ctx context.Context, db *mongo.Client) {
	err := godotenv.Load()
	if err != nil {
		log.Print("Failed to Parse the env guys")
		return
	}
	if ctx == nil || db == nil {
		log.Printf("Error during getting database contet and database db %s", ctx)
	}
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
	result, err := coll.InsertOne(ctx, user)
	if err != nil {
		log.Printf("Failed to insert %s", err.Error())
		_, err2 := writer.Write([]byte("Failed to insert"))
		if err2 != nil {
			return
		}
		os.Exit(0)
	}
	userId := result.InsertedID.(primitive.ObjectID)
	log.Printf("UserId of Inserted User is %s", userId.String())
	respBody := make(map[string]string)
	respBody["data"] = "User Successfully Created"
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Token", userId.String())
	err = json.NewEncoder(writer).Encode(&respBody)

	if err != nil {
		// Handle the error
		log.Printf(err.Error())
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (auth *AuthenticationHandler) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	log.Print(request.URL.Path)
	switch request.URL.Path {
	case "/lib/signup":
		Signup(write, request, auth.Ctx, auth.Db)
	}

}
