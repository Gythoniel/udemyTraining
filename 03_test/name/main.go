package main

import "fmt"

func main() {
	var name string
	fmt.Print("Please enter name: ")
	fmt.Scan(&name)
	fmt.Print("Hello, ", name, "!")
}
