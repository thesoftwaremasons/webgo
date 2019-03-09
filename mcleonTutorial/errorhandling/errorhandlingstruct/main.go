package main

import (
	"fmt"
	"log"
)

type ErrorStruct struct {
	lat  string
	long string
	err  error
}

func (e *ErrorStruct) Error() string {
	return fmt.Sprintf("some kind of error: %v %v %v", e.lat, e.long, e.err)
}
func main() {
	_, err := sqrt(-40)
	if err != nil {
		log.Println(err)
	}
}
func sqrt(f float64) (float64, error) {
	if f < 0 {
		err := fmt.Errorf("kinda of error")
		return 0, &ErrorStruct{"20", "40", err}
	}
	return 1, nil
}
