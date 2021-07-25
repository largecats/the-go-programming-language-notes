package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	url := os.Args[1]
	for i := 1; i < 3; i++ {
		outputPath := fmt.Sprintf(".\\exercise1.10\\output%d.txt", i)
		go fetch(url, ch, outputPath) // start a goroutine; calls fetch asynchronously
	}
	for i := 1; i < 3; i++ {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, outputFilePath string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		ch <- fmt.Sprintf("while creating output file %s: %v", outputFile.Name(), err)
		return
	}
	writer := bufio.NewWriter(outputFile)
	nbytes, err := io.Copy(writer, resp.Body) // write response to output file
	resp.Body.Close()                         // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
