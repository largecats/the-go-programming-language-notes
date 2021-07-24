package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:] // list of file names
	if len(files) == 0 {
		countLines(os.Stdin, counts) // if no file name is supplied, read from standard input
	} else {
		for _, arg := range files {
			f, err := os.Open(arg) // open the file using file name
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts) // count the number of appearances of the lines in the file and store in counts
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
