package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		io.WriteString(conn, "Hello world")
		fmt.Fprintln(conn, "How was your day")
		fmt.Fprintf(conn, "%v", "How was your day")
		conn.Close()
	}

}
