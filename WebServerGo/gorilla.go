package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func users(res http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userif := vars["id"]
	fmt.Fprintf(res, "User Id %s", userif)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", users).Methods("GET")
	http.ListenAndServe(":7777", r)
}
