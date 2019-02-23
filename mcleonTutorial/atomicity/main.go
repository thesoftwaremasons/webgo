package main

import (
	"fmt"
	"math/rand"

	"sync"
	"sync/atomic"
	"time"
)

var counter int64

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	fmt.Println("cool")
	go GetVal("Foo")
	go GetVal("Bar")
	wg.Wait()
	//run
	fmt.Println(counter)
}
func GetVal(s string) {

	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)

		//mutex.Lock()
		// counter++
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, counter)
		//mutex.Unlock()
	}
	wg.Done()
}
