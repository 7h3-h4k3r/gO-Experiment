package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"JwtAuthentication/authentication"
)

func protected_testcase(w http.ResponseWriter , r *http.Request){
	json.NewEncoder(w).Encode(map[string]string{"Test_case":"success authentiction workflow work awsome"})
}

func main() {
    http.HandleFunc("/login", authentication.SignIn)
    http.Handle("/protected", authentication.JwtMiddleware(http.HandlerFunc(protected_testcase)))

    fmt.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}