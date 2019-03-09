package main

import (
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		log.Println("Reason:", err)
	}
	log.SetOutput(nf)

}
func main() {
	_, err := os.Open("notext.txt")
	if err != nil {
		log.Println("Reason", err)
	}
}
