package authentication

import (
	"encoding/json"
	"io"
	"net/http"
)

type User struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignIn(w http.ResponseWriter , r *http.Request){
	
	if http.MethodPost != r.Method{
		http.Error(w,"Method Not allow to signin",http.StatusMethodNotAllowed)
		return
	}

	body , err := io.ReadAll(r.Body)
	if err!=nil{
		http.Error(w,"invalidate Request",http.StatusBadRequest)
		return
	}
	var in_user User

	err = json.Unmarshal(body,&in_user)

	if err!=nil || in_user.Username != "admin" || in_user.Password != "Pass@!23"{
		http.Error(w,"credentials has been Wrong",http.StatusNoContent)
		return
	}

	get_token ,err := GenToken(in_user.Username)

	if err!=nil{
		http.Error(w,"invalidate token Error :",http.StatusInternalServerError)
		return
	}


	json.NewEncoder(w).Encode(map[string]string{"token": get_token})

}