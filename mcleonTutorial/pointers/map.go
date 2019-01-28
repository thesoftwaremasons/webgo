package main

import "fmt"

func main() {
	allmaps := make(map[string]string, 5)
	allmaps["Time"] = "speek"
	allmaps["Ya"] = "Coo"
	//_, allmap = allmaps
	fmt.Println(allmaps["Ya"])

	aonther := map[int]string{
		0: "Zero",
		1: "One",
		2: "Two",
		3: "Three",
	}
	for k, v := range aonther {
		fmt.Println(k, v)
	}
	//fmt.Println(aonther)
}
