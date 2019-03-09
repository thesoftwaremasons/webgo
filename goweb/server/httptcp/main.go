package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	tp, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer tp.Close()
	for {
		conn, err := tp.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handler(conn)
	}
}
func handler(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()
	fmt.Println("I won't get here")
}
