package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	appearInFiles := make(map[string]map[string]bool) // line -> "set" of all files in which the line appears (implemented via map of filename -> true)
	files := os.Args[1:]                              // list of file names
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
			var filenames []string // extract filenames in which line appears
			for filename := range appearInFiles[line] {
				filenames = append(filenames, filename)
			}
			fmt.Printf("%d\t%s\t%v\n", n, line, filenames)
		}
	}
}

func countLines(f *os.File, counts map[string]int, appearInFiles map[string]map[string]bool, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if _, ok := appearInFiles[line]; ok { // if line is already present in appearInFiles
			if _, ok := appearInFiles[line][filename]; !ok { // add filename if it is not yet recorded
				appearInFiles[line][filename] = true
			}
		} else {
			appearInFiles[line] = map[string]bool{filename: true}
		}
	}
}

/*
> go run .\main.go
hello
hello
can you hear me
2       hello   [stdin stdin]

> go run .\main.go sample1.txt sample2.txt
3       hello   [sample1.txt sample2.txt]
*/
