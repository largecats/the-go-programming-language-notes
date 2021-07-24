package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] { // the first argument is the program name
		fmt.Println(index, arg)
	}
}
