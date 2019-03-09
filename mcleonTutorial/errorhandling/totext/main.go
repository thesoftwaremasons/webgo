package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("fe.txt")
	if err != nil {
		//fmt.Println("Why", err)
		log.Println("Reason: ", err)
		log.Println("Reason: ", err)
	}

}
