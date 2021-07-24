package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	appearInFiles := make(map[string][]string)
	files := os.Args[1:] // list of file names
	if len(files) == 0 {
		countLines(os.Stdin, counts, appearInFiles, "stdin") // if no file name is supplied, read from standard input
	} else {
		for _, arg := range files {
			f, err := os.Open(arg) // open the file using file name
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, appearInFiles, arg) // count the number of appearances of the lines in the file and store in counts
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, appearInFiles[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, appearInFiles map[string][]string, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		appearInFiles[line] = append(appearInFiles[line], filename)
	}
}

/*
> go run .\main.go
hello
hello
can you hear me
2       hello   [stdin stdin]

> go run .\main.go sample1.txt sample2.txt
2       hello   [sample1.txt sample1.txt]
*/
