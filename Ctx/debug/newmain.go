
package main

import _ "net/http/pprof"
import "net/http"

func main() {
    // Your app initialization here

    go func() {
        http.ListenAndServe("localhost:6060", nil)
    }()

    // your application logic
}