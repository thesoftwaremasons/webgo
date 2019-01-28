package main

import "fmt"

func main() {
	a := 2
	b := &a //b is getting the address
	fmt.Println(a, b)

	fmt.Println(*b)
	//change val with the pointer

	*b = 10
	fmt.Println(*b)
	fmt.Println(b)

}
