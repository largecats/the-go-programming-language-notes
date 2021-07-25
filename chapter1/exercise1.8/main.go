package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var PREFIX string = "http://"

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, PREFIX) {
			url = PREFIX + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body) // b for body
		resp.Body.Close()                   // need to close after reading
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

/*
...\chapter1> go build .\exercise1.8
...\chapter1> .\exercise1.8 gopl.io
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
          "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta name="go-import" content="gopl.io git https://github.com/adonovan/gopl.io"></meta>
  <title>The Go Programming Language</title>
...
...\chapter1> .\exercise1.8 bad.gopl.io
fetch: Get "http://bad.gopl.io": dial tcp: lookup bad.gopl.io: no such host
exit status 1
*/
