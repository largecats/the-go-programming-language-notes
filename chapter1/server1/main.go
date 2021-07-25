// Server1 is a minimal "echo" server
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)                         // connects the handler function to incoming URLs that begin with /
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) // starts a server listening for incoming requests on port 8000
}

// handler echoes the Path component of the request URL, r
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path) // sends response to ResponseWriter
}

/*
...\chapter1> go run .\server1\main.go # run in a separate window
*/

/*
...\chapter1> .\fetch http://localhost:8000 # or search in browser
URL.Path = "/"
...\chapter1> .\fetch http://localhost:8000/help
URL.Path = "/help"
*/
