package main

import "fmt"

func main() {
	c := fibonacci(30)
	fmt.Println(getfibonacci(c))
}

func fibonacci(n float64) chan float64 {
	c := make(chan float64)
	go func() {
		for i := n; i > 0; i-- {
			c <- i
		}
		close(c)
	}()
	return c
}
func getfibonacci(n chan float64) float64 {
	var score float64 = 1
	for i := range n {
		score *= i
	}
	return score
}
