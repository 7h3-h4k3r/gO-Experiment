package authentication

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type keys struct{
	Secret_key string `json:"secert_key"`
}

type Claim struct{
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func ReadEnv() string {
	file , err := os.ReadFile("../env")

	if err!=nil{
		panic(err)
	}

	var key keys;

	if err := json.Unmarshal(file,&key) ; err!=nil{
		panic(err);
	}
	return key.Secret_key
	
}

func GenToken(username string)(string , error){
	
	expireTime := time.Now().Add(1 * time.Hour)
	claims := Claim{
		Username:username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer: "authentication-creater-devops",
			Subject: username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return  token.SignedString([]byte(ReadEnv()))

}

// verify the Token base authentication 
func ItAuth(token_string string)(*Claim , error){

	claim := &Claim{}

	token , err := jwt.ParseWithClaims(token_string,claim,func(token *jwt.Token) (any, error) {

		if _ , ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return  nil , fmt.Errorf("un-authorized signing")
		}

		
		return  []byte(ReadEnv()), nil 
	})

	if err!=nil{
		fmt.Println("yes")
	}
	if err!=nil || !token.Valid{
		return nil, fmt.Errorf("invalidate Token")
	}

	return  claim , nil
}