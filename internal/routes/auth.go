package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/sangharshseth/internal/models"
	"github.com/sangharshseth/pkg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationHandler struct {
	Ctx context.Context
	Db  *mongo.Client
}

type UserDetails struct {
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Platform []string `json:"platform"`
}

type LoginDetails struct {
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
		log.Printf("Error during getting connections contet and connections db %s", ctx)
	}
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	var data UserDetails
	bodyParseError := decoder.Decode(&data)
	if bodyParseError != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.MinCost)
	user := models.User{
		Email:     data.Email,
		Password:  string(hashedPassword),
		Platform:  data.Platform,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
	coll := db.Database("development").Collection("Users")
	result, err := coll.InsertOne(ctx, user)
	headers := make(map[string]string)
	if err != nil {
		log.Printf("Failed to insert %s", err.Error())
		pkg.HTTPResponse("Failed to Create User", writer, http.StatusInternalServerError, headers)
	}
	userId := result.InsertedID.(primitive.ObjectID)
	log.Printf("UserId of Inserted User is %s", userId.String())
	writer.Header().Set("Token", userId.String())
	pkg.HTTPResponse("User Created Successfully", writer, http.StatusCreated, headers)
}

func Login(writer http.ResponseWriter, request *http.Request, ctx context.Context, db *mongo.Client) {
	err := godotenv.Load()
	if err != nil {
		log.Print("Failed to Parse the env guys")
		return
	}
	if ctx == nil || db == nil {
		log.Printf("Error during getting connections contet and connections db %s", ctx)
	}
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	var data LoginDetails
	bodyParseError := decoder.Decode(&data)
	if bodyParseError != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	UserCollection := db.Database("development").Collection("Users")
	var Result bson.M
	err = UserCollection.FindOne(ctx, bson.M{"email": data.Email}).Decode(&Result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Print("User Does not exists")
			http.Error(writer, "User Does not Exist", http.StatusBadRequest)
			return
		} else {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	hashedPw := Result["password"].(string)
	Headers := map[string]string{
		"Authorization": "Bearer Please Work",
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPw), []byte(data.Password))
	if err != nil {
		pkg.HTTPResponse("Authorization Error: Password Error", writer, http.StatusBadRequest, Headers)
		return
	}
	log.Print("Came before Response")
	pkg.HTTPResponse("Successfully Logged In, man", writer, http.StatusOK, Headers)

}

func (auth *AuthenticationHandler) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	log.Print(request.URL.Path)
	switch request.URL.Path {
	case "/auth/signup":
		Signup(write, request, auth.Ctx, auth.Db)
	case "/auth/login":
		Login(write, request, auth.Ctx, auth.Db)
	}

}
