package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string                   // implicitly initialized to empty strings
	for i := 1; i < len(os.Args); i++ { // for initialization; condition; post
		s += sep + os.Args[i]
		sep = " " // if we initialize sep := " " before the for loop, there'll be a blank space before the first argument (since os.Args[0], the program name, is not printed)
	}
	fmt.Println(s)
}
