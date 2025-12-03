package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type USER struct {
	ID   int `json:"id"`
	Name int `json:"name"`
}

func auth(w http.ResponseWriter, r *http.Request) {
	var user USER
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "ID Is %d Name is %d", user.ID, user.Name)
}
func main() {
	http.HandleFunc("/user", auth)
	http.ListenAndServe(":7070", nil)
}
