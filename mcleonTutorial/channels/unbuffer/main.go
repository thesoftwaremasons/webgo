package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		close(c)
	}()
	for n := range c {
		fmt.Println(n)
	}
	// go func() {
	// 	for {
	// 		fmt.Println(<-c)
	// 	}

	// }()
	time.Sleep(time.Second)
	//fmt.Print("ya")
}
