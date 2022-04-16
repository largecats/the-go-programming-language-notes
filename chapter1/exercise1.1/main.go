package main

import (
	"fmt"
)

func test(i int) (newI int) {
	newI = i
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("could not process: %d\n", i)
		}
	}()
	if i < 0 {
		panic("number is negative")
	}
	newI = i + 1
	return
}

func main() {
	// s, sep := "", ""
	// for _, arg := range os.Args {
	// 	s += sep + arg
	// 	sep = " "
	// }
	// fmt.Println(s)
	i := test(0)
	fmt.Printf("i = %d\n", i)
}
