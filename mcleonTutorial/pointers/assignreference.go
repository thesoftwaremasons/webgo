package main

import "fmt"

func main() {
	x := 40
	var y *int = &x
	fmt.Println(y)
}
