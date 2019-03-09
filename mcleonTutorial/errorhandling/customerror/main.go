package main

import (
	"errors"
	"log"
)

var errInfo = errors.New("Reason: no zero error")

func main() {
	_, err := sqrt(-32)
	if err != nil {
		log.Print(err)
	}

}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errInfo
	}
	return 42, nil
}
