package main

import "fmt"

func main() {
	fmt.Printf("")
	//map is key value pair
	//declare fitsy
	// email := make(map[string]string)

	// email["Bob"] = "Bob@gmail.com"
	// email["Sam"] = "sam@gmail.com"
	// //fmt.Print(email)
	// //fmt.Print(email["Bob"])
	// delete(email, "Bob")
	// fmt.Print(email["Bob"])

	//short way
	email := map[string]string{"Bob": "bob@gmail.com", "sharon": "sharon@gmail.com"}
	fmt.Println(email)

}
