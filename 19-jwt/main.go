package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	secret := []byte("INI_ADALAH_TOKEN_RAHASIA")

	claims := jwt.MapClaims {
		"uid" : "1234",
		"role" : "admin",
		"exp" : time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secret)

	if err != nil {
		panic(err)
	}

	fmt.Println("Generate JWT TOKEN:")
	fmt.Println(signedToken)
}