package main

import (
	"fmt"
	"math/rand"

	"sync"
	"time"
)

var counter int
var mutex sync.Mutex
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	fmt.Println("cool")
	go GetVal("Foo")
	go GetVal("Bar")
	wg.Wait()
	//
	fmt.Println(counter)
}
func GetVal(s string) {

	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)

		mutex.Lock()
		counter++
		fmt.Println(s, i, counter)
		mutex.Unlock()
	}
	wg.Done()
}
