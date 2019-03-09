package main

import "fmt"

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
	fmt.Println(temp)
}
