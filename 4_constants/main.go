package main

import "fmt"

// const age = 30

// var name = "Go"

func main() {
	// const pi = 3.14

	// fmt.Println(age)
	// fmt.Println(name)

	const (
		port = 8080
		host = "localhost"
	)
	fmt.Println(port, host)
}
