package routes

import (
	"CLI-Tools/auth"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type AuthenticationHandler struct{}

type UserDetails struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(writer http.ResponseWriter, request *http.Request) {
	err := godotenv.Load()
	if err != nil {
		return
	}
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	var data UserDetails
	err = decoder.Decode(&data)
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
	_, err = fmt.Fprint(writer, "SUCCESS", Jwt)
	if err != nil {
		return
	}
}

func DiscordOAuth2(writer http.ResponseWriter, request *http.Request) {
	var url = os.Getenv("DISCORD_GENERATED_URL")
	log.Printf("Generated URL is %s", url)
	http.Redirect(writer, request, url, 301)
}

func (auth *AuthenticationHandler) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	log.Print(request.URL.Path)
	switch request.URL.Path {
	case "/auth/signup":
		Signup(write, request)
	case "/auth/discord":
		log.Println("Came to discord oauth2")
		DiscordOAuth2(write, request)
	}

}
