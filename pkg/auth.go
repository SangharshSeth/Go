package pkg

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func VerifyAccess(name string) bool {
	if name == "Admin" {
		return true
	} else {
		return false
	}
}

func GenerateJWT(secret []byte, data string) (string, error) {
	JwtToken := jwt.NewWithClaims(
		jwt.SigningMethodHS512,
		jwt.MapClaims{
			"userdata": data,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		},
	)
	JwtTokenString, err := JwtToken.SignedString(secret)
	if err != nil {
		panic(err.Error())
	}
	return JwtTokenString, nil
}
