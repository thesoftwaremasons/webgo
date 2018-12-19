package main

import "fmt"

func main() {
	//fmt.Print("Cool")
	var fruitArr [2]string
	//Asssign Values
	fruitArr[0] = "Mango"
	fruitArr[1] = "Orange"

	//another way
	// fruit := [2]string{"Pineapple", "Apple"}

	// fmt.Println(fruit)

	//slice

	fruitSlice := []string{"age", "Or", "rage"}
	fmt.Println(fruitSlice)

}
