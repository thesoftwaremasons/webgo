package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)
	c1 := sq(in)
	c2 := sq(in)
	for c := range merge(c1, c2) {
		fmt.Println(c)
	}
}

func gen(nums ...int) chan int {
	c := make(chan int)
	go func() {
		for _, i := range nums {
			c <- i
		}
		close(c)
	}()
	return c
}
func sq(in chan int) chan int {
	c := make(chan int)
	go func() {
		for i := range in {
			c <- i * i
		}
		close(c)
	}()
	return c
}
func merge(in ...chan int) chan int {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(in))
	for _, i := range in {
		go func(ch chan int) {
			for v := range ch {
				c <- v
			}
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	return c
}
