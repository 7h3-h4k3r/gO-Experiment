package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func index(res http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(res, "Hello world From the Dharani Go web site")
}
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)

	loggedRouter := loggingMiddleware(r)

	fmt.Println("Server with logging middleware on :8080")
	http.ListenAndServe(":8080", loggedRouter)
}
