package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Femi"
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
