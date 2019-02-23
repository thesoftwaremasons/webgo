package main

import "fmt"

func main() {
	c := make(chan int)
	isDone := make(chan bool)

	go func() {
		for i := 0; i < 10000; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		isDone <- true
	}()
	go func() {
		for n := range c {
			fmt.Println(n)
		}
		isDone <- true
	}()
	<-isDone
	<-isDone
}
