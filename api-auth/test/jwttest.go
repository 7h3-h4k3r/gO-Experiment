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
	expiryTime := time.Now().Add(1 * time.Second)
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
func validateToken(token_s string)(*Clamis , error){

	claims := &Clamis{}

	token , err := jwt.ParseWithClaims(token_s,claims,func(token *jwt.Token)(interface{},error){
		if _ ,ok := token.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil , fmt.Errorf("unexpected signing method : %v ",token.Header["alg"])

		}
		return jwtkey,nil;
	})

	if err!= nil || !token.Valid{
		return nil,fmt.Errorf("invalid token : %v",err)
	}

	return claims,nil

}
func main(){

	// username := "dharani"
	// token ,err := genToken(username)
	// if err==nil{
	// 	fmt.Println(token)
	// }
	validate , err := validateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRoYXJhbmkiLCJpc3MiOiJhcGktdGVzdCIsInN1YiI6ImRoYXJhbmkiLCJleHAiOjE3NjczMjc4MzgsIm5iZiI6MTc2NzMyNzgzNywiaWF0IjoxNzY3MzI3ODM3fQ.KW8pD6FrQvlwYQRSvWIW9CY8W4gBk3px7Ck")
	if err==nil{
		fmt.Println(validate)
	}else{
		fmt.Println(validate ,err)
	}

	
	
	
}