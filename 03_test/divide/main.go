package main

import "fmt"

func main() {
	var small int
	var large int
	fmt.Print("enter small number: ")
	fmt.Scan(&small)
	fmt.Print("enter large number: ")
	fmt.Scan(&large)
	fmt.Print(large/small)
}
