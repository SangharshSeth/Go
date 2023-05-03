package routes

import (
	"CLI-Tools/auth"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

type AuthenticationHandler struct{}

type UserDetails struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(writer http.ResponseWriter, request *http.Request) {
	godotenv.Load()
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	var data UserDetails
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	env := os.Getenv("JWTSECRET")
	fmt.Println(env)
	Jwt, err := auth.GenerateJWT([]byte(env), data.Password)
	if err != nil {
		fmt.Println("Failed to generate JWT")
	}
	fmt.Fprint(writer, "SUCCESS", Jwt)
}

func (auth *AuthenticationHandler) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/signup":
		Signup(write, request)
	}
}
