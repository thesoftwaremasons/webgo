package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	name := os.Args[1]

	temp := `
	<html>
		<head>
			<title>Hello</title>
		</head>
		<body>
		<h3>` + name + `</h3>
		</body>
	</html>
	`
	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Reason is :", err)
	}
	defer file.Close()
	io.Copy(file, strings.NewReader(temp))

	fmt.Println(temp)
}
