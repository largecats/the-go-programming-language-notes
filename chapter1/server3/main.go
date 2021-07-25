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

// handler echoes the HTTP request
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

/*
...\chapter1> go run .\server3\main.go # in a separate window
*/

/*
...\chapter1> .\fetch http://localhost:8000
GET / HTTP/1.1
Header["User-Agent"] = ["Go-http-client/1.1"]
Header["Accept-Encoding"] = ["gzip"]
Host = "localhost:8000"
RemoteAddr = "127.0.0.1:7945"
*/
