// Server2 is a minimal "echo" and counter server
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() { // server runs the handler for each incoming request in a separate goroutine so that it can serve multiple requests simultaneously
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock() // ensure at most one goroutine can access the count variable at a time
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path - %q\n", r.URL.Path)
}

// counter echoes the number of calls so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

/*
...\chapter1> go run .\server2\main.go # in a separate window
*/

/*
...\chapter1> .\fetch http://localhost:8000
URL.Path - "/"
...\chapter1> .\fetch http://localhost:8000/count
Count 1
...\chapter1> .\fetch http://localhost:8000/help
URL.Path - "/help"
...\chapter1> .\fetch http://localhost:8000/count
Count 2
*/
