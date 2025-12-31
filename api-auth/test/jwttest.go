package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtkey = []byte("this-secret-key-stored-on-../env")

type Clamis struct{
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func genToken(username string) (string , error){
	expiryTime := time.Now().Add(1 * time.Hour)
	claims := Clamis{
		Username : username,
		RegisteredClaims : jwt.RegisteredClaims{
			ExpiresAt : jwt.NewNumericDate(expiryTime),
			IssuedAt : jwt.NewNumericDate(time.Now()),
			NotBefore : jwt.NewNumericDate(time.Now()),
			Issuer : "api-test",
			Subject : username,
		},

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(jwtkey)
}

func main(){

	username := "dharani"
	fmt.Println(genToken(username))
}