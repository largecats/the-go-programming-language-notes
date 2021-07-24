package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] { // the first argument is the program name
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
