package main

import "fmt"

func main() {
	fmt.Println("vook")

	//long method
	// i := 1
	// for i <= 10 {

	// 	fmt.Println(i)
	// 	i = i + 1
	// }
	// //short method
	// for i := 1; i <= 10; i++ {
	// 	println(i)
	// }

	//fizzbuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

	}
}
