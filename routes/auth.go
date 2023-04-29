package routes

import (
	"CLI-Tools/auth"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthenticationHandler struct{}

type UserDetails struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(writer http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	var data UserDetails
	err := decoder.Decode(&data)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	Jwt, err := auth.GenerateJWT([]byte("Sangharsh"), data.Password)
	_ = err

	fmt.Fprint(writer, "SUCCESS", Jwt)
}

func (auth *AuthenticationHandler) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/signup":
		Signup(write, request)
	}
}
