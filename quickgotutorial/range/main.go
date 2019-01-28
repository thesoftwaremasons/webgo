package main

import "fmt"

func main() {
	idos := map[string]string{"id": "1", "ide": "2", "ides": "3"}
	//arrays
	//ids := []int{32, 30, 10}
	// i := 0
	// for i <= len(ids) {
	// 	fmt.Println(ids[i])
	// 	i++
	// }

	//loop through range(foreach)

	// for _, id := range ids {
	// 	fmt.Println(id)

	// }

	for k, v := range idos {
		fmt.Println(k, v)
	}
	for _, v := range idos {
		fmt.Println("value :" + v)
	}
	//fmt.Println(ids[2:])
}
