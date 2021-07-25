package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		resp.Body.Close() // need to close after reading
	}
}

/*
...\chapter1> go build .\exercise1.7\
> .\exercise1.7 http://gopl.io
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
          "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta name="go-import" content="gopl.io git https://github.com/adonovan/gopl.io"></meta>
  <title>The Go Programming Language</title>
...
> .\exercise1.7 http://bad.gopl.io
fetch: Get "http://bad.gopl.io": dial tcp: lookup bad.gopl.io: no such host
exit status 1
*/
