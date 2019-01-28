package main

import "fmt"

func main() {

	//fmt.Println(display("kanu", "uas"))
	// n := displayTotal(1, 3, 5, 7, 9)
	// fmt.Println(n)
	// anono := func() {
	// 	fmt.Println("ya")
	// }
	// anono()
	//shoe := GreetMe()
	//fmt.Println(shoe())

	// tryCallBack([]int{1, 2, 3, 4}, func(n int) {
	// 	fmt.Println(n)
	// })
	// val := advanceCallback([]int{2, 4, 5, 5, 1, 3}, func(n int) bool {
	// 	if n > 1 {
	// 		return true
	// 	} else {
	// 		return false
	// 	}

	// })
	val := factorial(4)
	fmt.Println(val)
}
func display(firstname string, lastname string) (string, int) {
	return fmt.Sprint(firstname, lastname), len(firstname)
}
func displayTotal(sf ...float64) float64 {
	totalval := 0.0
	for _, i := range sf {
		totalval += i
	}
	return totalval
}
func GreetMe() func() string {
	return func() string {
		return "Ma"
	}
}

func tryCallBack(num []int, callback func(int)) {
	for _, number := range num {
		callback(number)
	}
}

func advanceCallback(num []int, newcallback func(int) bool) []int {
	//declare empty array
	emp := []int{}
	for _, n := range num {
		if newcallback(n) {
			emp = append(emp, n)
		}
	}
	return emp
}
func factorial(x int) int {
	if x == 0 {
		return 1
	} else {
		return x * factorial(x-1)
	}
}
