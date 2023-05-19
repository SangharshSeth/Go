package routes

import (
	"CLI-Tools/auth"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/url"
	"os"
)

type AuthenticationHandler struct{}

type UserDetails struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type OAuthResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func Signup(writer http.ResponseWriter, request *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Print("Failed to Parse the env guys")
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
	var DiscordForwardURL = os.Getenv("DISCORD_GENERATED_URL")
	log.Printf("Generated URL is %s", DiscordForwardURL)
	http.Redirect(writer, request, DiscordForwardURL, 301)
}

func HandleDiscordOAuth2Callback(writer http.ResponseWriter, request *http.Request) {
	var code = request.URL.Query().Get("code")
	_, err := fmt.Fprint(writer, code)
	if err != nil {
		log.Panic("Error is", err.Error())
	}

	httpClient := &http.Client{}

	var body = url.Values{}
	body.Set("client_id", os.Getenv("DISCORD_CLIENT_ID"))
	body.Set("client_secret", os.Getenv("DISCORD_CLIENT_SECRET"))
	body.Set("grant_type", "authorization_code")
	body.Set("code", code)
	body.Set("redirect_uri", os.Getenv("DISCORD_REDIRECT_URL"))

	err = godotenv.Load()

	httpRequest, err := http.NewRequest("POST", "https://discord.com/api/oauth2/token", bytes.NewBufferString(body.Encode()))
	httpRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	httpRequest.Header.Set("Accept-Encoding", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(httpRequest)

	decoded := json.NewDecoder(resp.Body)

	var OAuthData OAuthResponse
	decodeErr := decoded.Decode(&OAuthData)
	if decodeErr != nil {
		log.Print("Failed to decode OAuth2 data", decodeErr.Error())
	}

	var store = sessions.NewCookieStore([]byte(os.Getenv("SESSIONKEY")))

	session, _ := store.Get(request, "session-name")
	session.Values["authToken"] = OAuthData.AccessToken
	err = session.Save(request, writer)
	http.Redirect(writer, request, os.Getenv("DISCORD_REDIRECT_URL"), http.StatusSeeOther)
}

func (auth *AuthenticationHandler) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	log.Print(request.URL.Path)
	switch request.URL.Path {
	case "/auth/signup":
		Signup(write, request)
	case "/auth/discord":
		DiscordOAuth2(write, request)
	case "/auth/discord_callback":
		HandleDiscordOAuth2Callback(write, request)
	}

}
