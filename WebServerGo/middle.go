func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloHandler)

	loggedRouter := loggingMiddleware(r)

	fmt.Println("Server with logging middleware on :8080")
	http.ListenAndServe(":8080", loggedRouter)
}