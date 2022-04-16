package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	filenames := []string{"1001", "1002", "1003"}
	res := makeThumbnails(filenames)
	//    res := makeThumbnails3(filenames)
	fmt.Println(res)
}

func makeThumbnails(filenames []string) int {
	sizes := make(chan int)
	var wg sync.WaitGroup

	for _, f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			size, _ := strconv.Atoi(f)
			sizes <- size
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int
	for size := range sizes {
		total += size
	}

	return total
}

func makeThumbnails1(filenames []string) int {
	sizes := make(chan int, 3)
	var wg sync.WaitGroup

	for _, f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			size, _ := strconv.Atoi(f)
			sizes <- size
		}(f)
	}

	// closer - deadlock if sizes is unbuffered
	wg.Wait()
	close(sizes)

	var total int
	for size := range sizes {
		total += size
	}

	return total
}

func makeThumbnails2(filenames []string) int {
	sizes := make(chan int, 3)
	var wg sync.WaitGroup

	for _, f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			size, _ := strconv.Atoi(f)
			sizes <- size
		}(f)
	}

	// closer - deadlock if sizes is unbuffered
	wg.Wait()
	close(sizes)

	var total int
	if len(sizes) > 0 { // if the wait and close are in another goroutine, this block will be skipped because len(sizes) will be 0 when the flow reaches here
		for size := range sizes {
			total += size
		}
	}

	return total
}

func makeThumbnails3(filenames []string) int {
	sizes := make(chan int, 3)
	var wg sync.WaitGroup

	for _, f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			size, _ := strconv.Atoi(f)
			sizes <- size
		}(f)
	}

	// closer - deadlock if sizes is unbuffered
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int
	if len(sizes) > 0 { // if the wait and close are in another goroutine, this block will be skipped because len(sizes) will be 0 when the flow reaches here
		for size := range sizes {
			total += size
		}
	}

	return total
}
