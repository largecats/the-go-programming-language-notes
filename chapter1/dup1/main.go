package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // map of line -> number of appearances
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { // reads the next line without \n
		counts[input.Text()]++
	}
	for line, n := range counts { // print each line that appears more than once; unordered
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

/*
hello
hello
can you hear me
2       hello
*/
