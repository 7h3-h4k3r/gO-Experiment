package main

import (
	"fmt"
	"net/http"
)

func index(res http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(res, "Hello world From the Dharani Go web site")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Print("Service Start with 4545 Port Number")
	if err := http.ListenAndServe(":4545", nil); err != nil {
		fmt.Printf("Service suddenly stoped somewent is Wrong")
	}

}
