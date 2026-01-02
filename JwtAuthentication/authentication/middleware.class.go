package authentication

import (
	"net/http"
	"strings"
	"fmt"
)



func JwtMiddleware(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth_header := r.Header.Get("Authorization")

		if auth_header == ""{
			http.Error(w,"Authorization header Missing",http.StatusUnauthorized)
			return 
		}

		token := strings.Split(auth_header, " ")

		if len(token)!=2 || strings.ToLower(token[0]) != "bearer"{
			http.Error(w,"invalid authotization data",http.StatusUnauthorized)
			return 
		}
		fmt.Println(token[1])
		if _ , err := ItAuth(token[1]);err!=nil{
			http.Error(w,"invalid token :"+err.Error(),http.StatusUnauthorized)
			return 
		}
		
		next.ServeHTTP(w,r)
	})
}