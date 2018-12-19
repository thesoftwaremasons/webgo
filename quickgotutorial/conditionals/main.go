package main

import "fmt"

func main() {
	fmt.Println("ya")
	x, y := 15, 10
	if x <= y {
		fmt.Printf("%d is less than %d", x, y)
	} else {
		fmt.Printf("%d is less than %d", y, x)
	}

	color := ""
	if color == "red" {
		fmt.Printf("color is red")
	} else if color == "blue" {
		fmt.Print("color is blue")
	} else {
		fmt.Print("Error Happened")
	}

	switch color {
	case "red":
		fmt.Print("is red")
	case "blue":
		fmt.Print("is blue")
	default:
		fmt.Print("Err")
	}

	//conditionals
}
