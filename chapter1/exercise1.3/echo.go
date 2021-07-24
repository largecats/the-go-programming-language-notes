package echo

import (
	"fmt"
	"strings"
)

func echo1(args []string) {
	s, sep := "", ""
	for _, arg := range args { // the first argument is the program name
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3(args []string) {
	fmt.Println(strings.Join(args, " "))
}
