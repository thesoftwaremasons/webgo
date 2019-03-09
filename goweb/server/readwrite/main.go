package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()
	for {

		connaccepted, err := listener.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handler(connaccepted)
	}

}
func handler(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(time.Second * 10))
	if err != nil {
		log.Panic("Connection Timeout")
	}
	readConn := bufio.NewScanner(conn)
	for readConn.Scan() {
		readline := readConn.Text()
		fmt.Println(readline)
		fmt.Fprintf(conn, "Hello femi :%s\n:", readline)
	}

	conn.Close()
	fmt.Print("Code got here")
}
