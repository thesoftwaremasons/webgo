package main

import (
	"fmt"
	"sync"
)

func main() {
	in := fanner()
	c1 := factorial(in)
	c2 := factorial(in)
	c3 := factorial(in)
	c4 := factorial(in)
	c5 := factorial(in)

	for i := range merge(c1, c2, c3, c4, c5) {
		fmt.Println(i)
	}
	// c := fanner(100)
	// fmt.Println(implementor(c))
}

//fan in
//merge all channels into 1

// func fanner(n float64) <-chan float64 {
// 	channel := make(chan float64)
// 	go func() {
// 		for i := n; i > 0; i-- {
// 			channel <- i
// 		}
// 		close(channel)
// 	}()
// 	return channel
// }
// func implementor(n <-chan float64) float64 {
// 	//channel := make(chan float64)
// 	var score float64 = 1
// 	for i := range n {
// 		score *= i
// 	}
// 	return score
// }
func fanner() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			for j := 3; i < 13; j++ {
				c <- j
			}
		}
		close(c)
	}()
	return c
}
func factorial(n <-chan int) <-chan int {
	channel := make(chan int)
	go func() {
		for i := range n {
			channel <- fact(i)
		}
		close(channel)
	}()
	return channel
}
func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
func merge(n ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	channel := make(chan int)
	output := func(c <-chan int) {
		for n := range c {
			channel <- n
		}
		wg.Done()
	}
	wg.Add(len(n))
	for _, c := range n {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(channel)
	}()
	return channel
}
