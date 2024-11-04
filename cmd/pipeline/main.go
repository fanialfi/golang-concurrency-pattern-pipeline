package main

import (
	"log"
	"sync"
)

func main() {
	ch := gen(1, 2, 3, 4, 5, 6, 7, 8, 9)

	// mendistribusikan square task ke 3 goroutine yang ketiganya membaca dari ch
	out1 := sq(ch)
	out2 := sq(ch)
	out3 := sq(ch)

	// convert list of channels to single channel by starting for each inbound channel
	out := merge(out1, out2, out3)

	for dataInt := range out {
		log.Println(dataInt)
	}
}

func gen(numbs ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, val := range numbs {
			out <- val
		}
		close(out)
	}()

	return out
}

// this function squaring data from chanIn and return squared number as channel
func sq(chanIn <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range chanIn {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func merge(manyChanIn ...<-chan int) <-chan int {
	out := make(chan int)
	wg := new(sync.WaitGroup)
	process := func(chanIn <-chan int) {
		for chanInt := range chanIn {
			out <- chanInt
		}
		wg.Done()
	}

	wg.Add(len(manyChanIn))
	for _, chanIn := range manyChanIn {
		go process(chanIn)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
