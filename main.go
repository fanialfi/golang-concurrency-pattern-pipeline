package main

import "fmt"

func main() {
	// set up the pipeline
	c := gen(1, 2, 3, 4, 5)
	out := square(c)

	// mengkonsumsi output
	for n := range out {
		fmt.Println(n)
	}

}

// sebuah function yang men-convert sebuah list of integer kedalam channel
// yang mana akan mengeluarkan (emit) bilangan-bilangan bulat dari list satu persatu
func gen(numbs ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range numbs {
			out <- n
		}
		close(out)
	}()

	return out
}

// menerima sebuah integer dari channel dan mereturn sebuah channel yang mengeluarkan (emit)
// channel yang telah di square (dipangkatkan) untuk setiap integer yang telah diterima dari channel
func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}

		close(out)
	}()

	return out
}
