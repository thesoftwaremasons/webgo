package main

import "fmt"

func main() {
	c := incrementor()
	output(c)
}
func incrementor() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func output(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
}
